package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// generated by protoc-gen-go
	"go.sazak.io/proxyqueue/gen/proxyqueue/v1/proxyqueuev1connect" // generated by protoc-gen-connect-go
	"go.sazak.io/proxyqueue/internal/server"
)

var (
	configPath = flag.String("config", "config.json", "path to the config file")
	listenAddr = flag.String("listen", ":8080", "address to listen on")
)

func main() {
	flag.Parse()

	srv, err := server.New(*configPath)
	if err != nil {
		log.Fatalf("Failed to create server: %s", err)
	}
	mux := http.NewServeMux()
	path, handler := proxyqueuev1connect.NewProxyQueueServiceHandler(srv)
	mux.Handle(path, handler)

	log.Printf("Listening on %s", *listenAddr)

	http.ListenAndServe(
		*listenAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
