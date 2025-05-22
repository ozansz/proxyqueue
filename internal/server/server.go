package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"connectrpc.com/connect"
	proxyqueuev1 "go.sazak.io/proxyqueue/gen/proxyqueue/v1"
	"go.sazak.io/proxyqueue/gen/proxyqueue/v1/proxyqueuev1connect"
	"go.sazak.io/proxyqueue/internal/storage"
	njq "go.sazak.io/proxyqueue/pkg/notifyingjobqueue"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	cacheTTL = 1 * time.Hour
)

type Server interface {
	proxyqueuev1connect.ProxyQueueServiceHandler

	Close() error
}

type server struct {
	storage storage.Storage[string, *proxyqueuev1.SubmitURLResponse]
	queue   njq.NotifyingJobQueue[*proxyqueuev1.SubmitURLResponse]

	configLock sync.Mutex
	config     *Config
	proxyIndex int
}

func New(configPath string) (Server, error) {
	config, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	return &server{
		config:  config,
		storage: storage.NewMemoryStorage[string, *proxyqueuev1.SubmitURLResponse](),
		queue:   njq.NewNotifyingJobQueue[*proxyqueuev1.SubmitURLResponse](config.Concurrency),
	}, nil
}

func (s *server) SubmitURL(
	ctx context.Context,
	req *connect.Request[proxyqueuev1.SubmitURLRequest],
) (*connect.Response[proxyqueuev1.SubmitURLResponse], error) {
	s.configLock.Lock()
	proxyIndex := s.proxyIndex
	proxy := s.config.Proxies[proxyIndex]
	s.proxyIndex = (proxyIndex + 1) % len(s.config.Proxies)
	s.configLock.Unlock()

	createdAt := timestamppb.New(time.Now())

	ch := s.queue.Submit(ctx, func(workerID int) (*proxyqueuev1.SubmitURLResponse, error) {
		cacheKey := "plain/" + req.Msg.Url
		if req.Msg.BinaryContent {
			cacheKey = "binary/" + req.Msg.Url
		}

		record, err := s.storage.Get(cacheKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get record: %w", err)
		}
		if record != nil {
			log.Printf("Worker %d found record for %s in cache", workerID, req.Msg.Url)
			return record, nil
		}

		startTime := timestamppb.New(time.Now())
		log.Printf("Worker %d is processing request for %s", workerID, req.Msg.Url)

		proxyURL, err := url.Parse(proxy)
		if err != nil {
			log.Printf("Worker %d failed to parse proxy: %s", workerID, err)
			return nil, fmt.Errorf("invalid proxy: %w", err)
		}
		cl := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				ForceAttemptHTTP2: false,
			},
			Timeout: 15 * time.Second,
		}

		r, err := http.NewRequestWithContext(ctx, http.MethodGet, req.Msg.Url, nil)
		if err != nil {
			log.Printf("Worker %d failed to create request: %s", workerID, err)
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		if req.Msg.UserAgent != nil {
			r.Header.Set("User-Agent", *req.Msg.UserAgent)
		} else if req.Msg.UseRandomUserAgent {
			r.Header.Set("User-Agent", s.config.UserAgents.Random())
		}

		for key, value := range req.Msg.Headers {
			r.Header.Set(key, value)
		}

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("Worker %d failed to dump request: %s", workerID, err)
			return nil, fmt.Errorf("failed to dump request: %w", err)
		}
		log.Printf("Worker %d dumped request: %s", workerID, string(dump))
		log.Printf("Worker %d sending request to %s with headers %v and proxy %s", workerID, req.Msg.Url, r.Header, proxy)

		resp, err := cl.Do(r)
		if err != nil {
			log.Printf("Worker %d failed to send request: %s", workerID, err)
			return nil, fmt.Errorf("failed to send request: %w", err)
		}
		log.Printf("Worker %d received response with status code %d", workerID, resp.StatusCode)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Worker %d failed to read response body: %s", workerID, err)
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		log.Printf("Worker %d read response body with length %d", workerID, len(body))
		finishedAt := timestamppb.New(time.Now())
		log.Printf("Worker %d finished processing request for %s in %.2f seconds (%.2f seconds idle time) with status code %d", workerID, req.Msg.Url, time.Since(startTime.AsTime()).Seconds(), time.Since(createdAt.AsTime()).Seconds()-time.Since(startTime.AsTime()).Seconds(), resp.StatusCode)

		record = &proxyqueuev1.SubmitURLResponse{
			Url:        req.Msg.Url,
			CreatedAt:  createdAt,
			StartedAt:  startTime,
			FinishedAt: finishedAt,
		}
		if req.Msg.BinaryContent {
			record.BinaryContent = body
		} else {
			record.HtmlContent = string(body)
		}

		log.Printf("Worker %d setting record for %s in cache", workerID, req.Msg.Url)
		s.storage.Set(cacheKey, record, cacheTTL)
		log.Printf("Worker %d set record for %s in cache", workerID, req.Msg.Url)

		return record, nil
	})

	res := <-ch
	if res.Err != nil {
		return nil, connect.NewError(connect.CodeInternal, res.Err)
	}
	return connect.NewResponse(res.Res), nil
}

func (s *server) Close() error {
	return s.queue.Close()
}
