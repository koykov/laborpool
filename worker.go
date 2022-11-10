package laborpool

import (
	"context"
	"sync/atomic"
)

type Worker struct {
	status uint32
	jobs   chan JobFn
	ctx    context.Context
	cancel context.CancelFunc
}

func trainWorker() *Worker {
	w := Worker{
		status: 1,
		jobs:   make(chan JobFn, 1),
	}
	w.ctx, w.cancel = context.WithCancel(context.Background())
	go w.wait(w.ctx)
	return &w
}

func (w *Worker) wait(ctx context.Context) {
	for {
		select {
		case job := <-w.jobs:
			_ = job()
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) Do(job JobFn) error {
	if !w.checkStatus() {
		return ErrWorkerStatus
	}
	w.jobs <- job
	return nil
}

func (w *Worker) Release() {
	if !w.checkStatus() {
		return
	}
	w.cancel()
	close(w.jobs)
}

func (w *Worker) checkStatus() bool {
	return atomic.LoadUint32(&w.status) == 1
}
