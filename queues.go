package multipleJobQueue

import (
	"fmt"
	"reflect"
)

var internalQueues *queues

type option func(*queues)

const defaultQueue = "default"

type queues struct {
	queues  map[string]*queue
	numJobs int
}

// WaitUntilFinish waits until all jobs has been executed
func (s *queues) WaitUntilFinish() {
	for {
		if s.numJobs == 0 {
			break
		}
	}
}

func multipleSelect(chans []chan Job) (int, Job, bool) {
	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}

	ch, v, ok := reflect.Select(cases)
	return ch, v.Interface().(Job), ok
}

func (s *queues) dispatch() {

	var queues []chan Job
	var queueKey []string
	for key := range s.queues {
		queues = append(queues, s.queues[key].queue)
		queueKey = append(queueKey, key)
	}
	for {
		if ch, job, ok := multipleSelect(queues); ok {

			if ok {
				queue := s.queues[queueKey[ch]]
				go func() {
					workerJobQueue := <-queue.pool
					workerJobQueue <- job
				}()
			}
		}
	}
}

func (s *queues) getQueue(alias string) (*queue, error) {
	queue, ok := s.queues[alias]
	if !ok {
		return nil, fmt.Errorf("queue not found")
	}
	return queue, nil
}

// AddQueue adds a new queue with a number of workers
func AddQueue(alias string, numWorkers int) option {
	return func(s *queues) {
		queue := createQueue(alias, numWorkers)
		s.queues[alias] = queue
	}
}

// InitializeQueues creates a default queue with a number of workers and dispatch each queue
func InitializeQueues(numWorkersInDefaultQueue int, options ...option) *queues {
	internalQueues = &queues{
		queues:  make(map[string]*queue),
		numJobs: 0,
	}

	queue := createQueue(defaultQueue, numWorkersInDefaultQueue)
	internalQueues.queues[defaultQueue] = queue

	for _, option := range options {
		option(internalQueues)
	}

	go internalQueues.dispatch()

	return internalQueues
}
