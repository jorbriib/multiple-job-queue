package multiple_job_queue

type Job interface {
	Handle()
}