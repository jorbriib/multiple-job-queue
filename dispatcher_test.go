package multiple_job_queue

import (
	"testing"
)

func TestGetDispatcher(t *testing.T) {
	dispatcher := GetDispatcher()
	if dispatcher == nil {
		t.Errorf("dispatcher is nil")
	}
}

func TestDispatcher_Dispatch_ReturnErrorIfQueuesWereNotInitialized(t *testing.T) {
	job := &TestJob{}

	dispatcher := GetDispatcher()
	err := dispatcher.Dispatch(job)
	if err == nil {
		t.Errorf("error is not nill when queues were not initialized")
	}
}

func TestDispatcher_Dispatch_ReturnErrorIfQueueNotExist(t *testing.T) {
	InitializeQueues(0)

	job := &TestJob{}

	dispatcher := GetDispatcher()
	err := dispatcher.Dispatch(job, "no_existing_queue")
	if err == nil {
		t.Errorf("error is not nill when queue_not_exists")
	}
}

func TestDispatcher_Dispatch(t *testing.T) {
	queues := InitializeQueues(0)

	job := &TestJob{}

	go func(t *testing.T) {
		for {
			select {
			case _ = <-queues.queues[DefaultQueue].queue:
				return

			default:
				t.Errorf("job not enqueued")
			}
		}
	}(t)

	dispatcher := GetDispatcher()
	err := dispatcher.Dispatch(job, DefaultQueue)
	if err != nil {
		t.Errorf(err.Error())
	}
	if queues.numJobs != 1 {
		t.Errorf( "numJobs property was not updated")
	}
}