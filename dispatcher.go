package multiple_job_queue

import "fmt"

type Dispatcher struct {
}

func (d *Dispatcher) Dispatch(job Job, aliasQueue ...string) error {
	if queues == nil {
		return fmt.Errorf("queues are not initialized")
	}

	selectedQueue := DefaultQueue
	if len(aliasQueue) > 0 {
		selectedQueue = aliasQueue[0]
	}

	queue, err := queues.getQueue(selectedQueue)
	if err != nil {
		return err
	}

	queue.enqueue(job)
	queues.numJobs++

	return nil
}

func GetDispatcher() *Dispatcher {
	return &Dispatcher{}
}
