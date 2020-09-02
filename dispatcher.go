package multiplejobqueue

import "fmt"

type dispatcher struct {
}

// Dispatch dispatches a job to a queue, by default the job is dispatched to the default queue
func (d *dispatcher) Dispatch(job Job, aliasQueue ...string) error {
	if internalQueues == nil {
		return fmt.Errorf("internalQueues are not initialized")
	}

	selectedQueue := defaultQueue
	if len(aliasQueue) > 0 {
		selectedQueue = aliasQueue[0]
	}

	queue, err := internalQueues.getQueue(selectedQueue)
	if err != nil {
		return err
	}

	queue.enqueue(job)
	internalQueues.numJobs++

	return nil
}

// GetDispatcher returns a new dispatcher to dispatch jobs
func GetDispatcher() *dispatcher {
	return &dispatcher{}
}
