package multiple_job_queue

import (
	"testing"
)

type TestJob struct {
}

func (j *TestJob) Handle() {
}

func TestQueue_Enqueue(t *testing.T) {

	job := &TestJob{}

	jobQueue := &Queue{
		alias:   "alias",
		queue:   make(chan Job),
		workers: make([]*Worker, 0),
		pool:    make(chan chan Job, 0),
	}

	go func(t *testing.T) {
		for {
			select {
			case _ = <-jobQueue.queue:
				return

			default:
				t.Errorf("job not enqueued")
			}
		}
	}(t)
	jobQueue.enqueue(job)
}
