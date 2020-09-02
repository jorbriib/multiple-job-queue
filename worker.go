package multiple_job_queue


type Worker struct {
	queue    chan Job
	pool     chan chan Job
}

func (w *Worker) start() {
	go func() {
		for {
			w.pool <- w.queue

			select {
			case job := <-w.queue:
				job.Handle()
				queues.numJobs--
			}
		}
	}()
}

func newWorker(pool chan chan Job) *Worker {
	return &Worker{
		queue:    make(chan Job),
		pool:     pool,
	}
}
