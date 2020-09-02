package multiplejobqueue

type queue struct {
	alias   string
	queue   chan Job
	workers []*worker
	pool    chan chan Job
}

func (q *queue) enqueue(job Job) {
	q.queue <- job
}

func createQueue(alias string, numWorkers int) *queue {
	jobQueue := &queue{
		alias:   alias,
		queue:   make(chan Job),
		workers: make([]*worker, numWorkers),
		pool:    make(chan chan Job, numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		jobQueue.workers[i] = newWorker(jobQueue.pool)
		jobQueue.workers[i].start()
	}

	return jobQueue
}
