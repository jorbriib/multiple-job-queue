package multipleJobQueue

type worker struct {
	queue chan Job
	pool  chan chan Job
}

func (w *worker) start() {
	go func() {
		for {
			w.pool <- w.queue

			select {
			case job := <-w.queue:
				job.Handle()
				internalQueues.numJobs--
			}
		}
	}()
}

func newWorker(pool chan chan Job) *worker {
	return &worker{
		queue: make(chan Job),
		pool:  pool,
	}
}
