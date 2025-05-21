package njq

import (
	"context"
	"sync"
)

const (
	bufferSize = 10000
)

type Doer[T any] func(workerID int) (T, error)

type Job[T any] struct {
	do         Doer[T]
	resultChan chan<- *JobResult[T]
}

type JobResult[T any] struct {
	Err error
	Res T
}

type NotifyingJobQueue[T any] interface {
	Submit(context.Context, Doer[T]) <-chan *JobResult[T]
	Close() error
}

type notifyingJobQueue[T any] struct {
	jobs   chan *Job[T]
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

type worker[T any] struct {
	id   int
	jobs <-chan *Job[T]
	wg   *sync.WaitGroup
}

func NewNotifyingJobQueue[T any](numWorkers int) NotifyingJobQueue[T] {
	ctx, cancel := context.WithCancel(context.Background())

	q := &notifyingJobQueue[T]{
		jobs:   make(chan *Job[T], bufferSize),
		cancel: cancel,
		wg:     sync.WaitGroup{},
	}

	for i := range numWorkers {
		w := &worker[T]{
			id:   i,
			jobs: q.jobs,
			wg:   &q.wg,
		}
		go w.work(ctx)
	}

	return q
}

func (w *worker[T]) work(ctx context.Context) {
	w.wg.Add(1)
	defer w.wg.Done()

	_work := func(job *Job[T]) {
		res, err := job.do(w.id)
		job.resultChan <- &JobResult[T]{Res: res, Err: err}
		close(job.resultChan)
	}

	for {
		select {
		case <-ctx.Done():
			// Drain the channel
			for job := range w.jobs {
				_work(job)
			}
			return
		case job := <-w.jobs:
			_work(job)
		}
	}
}

func (q *notifyingJobQueue[T]) Submit(ctx context.Context, do Doer[T]) <-chan *JobResult[T] {
	ch := make(chan *JobResult[T])
	job := &Job[T]{
		do:         do,
		resultChan: ch,
	}
	q.jobs <- job
	return ch
}

func (q *notifyingJobQueue[T]) Close() error {
	q.cancel()
	close(q.jobs)
	q.wg.Wait()
	return nil
}
