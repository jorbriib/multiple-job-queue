package multipleJobQueue

// Job is an interface, which all jobs should implement
type Job interface {
	Handle()
}