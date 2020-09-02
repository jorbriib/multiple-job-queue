package multipleJobQueue

import (
	"testing"
)

func TestNewWorker(t *testing.T) {
	worker := newWorker(make(chan chan Job))
	if worker == nil {
		t.Errorf("worker is nil")
	}
}

func TestWorker_Start(t *testing.T) {
	queues := InitializeQueues(0)

	pool := make(chan chan Job)
	worker := newWorker(pool)

	go func(t *testing.T) {
		for {
			select {
			case _ = <-worker.pool:
				return

			default:
				t.Errorf("job not enqueued")
			}
		}
	}(t)

	worker.start()

	job := &TestJob{}
	worker.queue <- job

	if queues.numJobs > 0 {
		t.Errorf("job not handled")
	}
}