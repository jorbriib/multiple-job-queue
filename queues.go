package multiple_job_queue

import (
	"fmt"
	"reflect"
)

var queues *Queues

type Option func(*Queues)

const DefaultQueue = "default"

type Queues struct {
	queues  map[string]*Queue
	numJobs int
}


func (s *Queues) WaitUntilFinish() {
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

func (s *Queues) dispatch() {

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


func (s *Queues) getQueue(alias string) (*Queue, error) {
	queue, ok := s.queues[alias]
	if !ok {
		return nil, fmt.Errorf("queue not found")
	}
	return queue, nil
}

func AddQueue(alias string, numWorkers int) Option {
	return func(s *Queues) {
		queue := createQueue(alias, numWorkers)
		s.queues[alias] = queue
	}
}

func InitializeQueues(numWorkersInDefaultQueue int, options ...Option) *Queues {
	queues = &Queues{
		queues:  make(map[string]*Queue),
		numJobs: 0,
	}

	queue := createQueue(DefaultQueue, numWorkersInDefaultQueue)
	queues.queues[DefaultQueue] = queue

	for _, option := range options {
		option(queues)
	}

	go queues.dispatch()

	return queues
}