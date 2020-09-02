package multiple_job_queue


type Queue struct {
	alias   string
	queue   chan Job
	workers []*Worker
	pool    chan chan Job
}

func (q *Queue) enqueue(job Job) {
	q.queue <- job
}

func createQueue(alias string, numWorkers int) *Queue {
	jobQueue := &Queue{
		alias:   alias,
		queue:   make(chan Job),
		workers: make([]*Worker, numWorkers),
		pool:    make(chan chan Job, numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		jobQueue.workers[i] = newWorker(jobQueue.pool)
		jobQueue.workers[i].start()
	}

	return jobQueue
}
