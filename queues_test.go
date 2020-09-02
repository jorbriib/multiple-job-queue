package multiplejobqueue

import "testing"

func TestInitializeQueues_CreatesADefaultQueue(t *testing.T) {
	queues := InitializeQueues(3)
	if queues.numJobs != 0 {
		t.Errorf("intialize internalQueues has a wrong numJobs value")
	}
	if len(queues.queues) != 1 {
		t.Errorf("intialize internalQueues doesn't create the default queue")
	}
	if len(queues.queues[defaultQueue].workers) != 3 {
		t.Errorf("initialize queue doesn't start three workers")
	}
	if queues.queues[defaultQueue].alias != defaultQueue {
		t.Errorf("wrong alias")
	}
}

func TestAddQueue_CreatesANewQueue(t *testing.T) {
	queues := InitializeQueues(3,
		AddQueue("high", 5))
	if len(queues.queues) != 2 {
		t.Errorf("intialize internalQueues doesn't create the default queue")
	}
	if len(queues.queues[defaultQueue].workers) != 3 {
		t.Errorf("initialize queue doesn't start three workers")
	}
	if queues.queues[defaultQueue].alias != defaultQueue {
		t.Errorf("wrong alias")
	}
	if len(queues.queues["high"].workers) != 5 {
		t.Errorf("initialize queue doesn't start five workers in high queue")
	}
	if queues.queues["high"].alias != "high" {
		t.Errorf("wrong alias")
	}
}
