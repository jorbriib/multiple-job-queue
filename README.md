[![Go Report Card](https://goreportcard.com/badge/github.com/jorbriib/multiple-job-queue)](https://goreportcard.com/report/github.com/jorbriib/multiple-job-queue) 
[![Build Status](https://travis-ci.com/jorbriib/multiple-job-queue.svg?branch=master)](https://travis-ci.com/jorbriib/multiple-job-queue)

Multiple Job Queue provides an API to process jobs in asynchronous way. 
Also, it allows you to defer the processing of a time-consuming task 
and creates a number of workers to parallelize the execution of each job.

Installation
------------

The routing package is just a common Go module. You can install it as any other Go module. 
```shell script
go get github.com/jorbriib/multiple-job-queue
```
To get more information just review the official [Go blog][1] regarding this topic.

Usage
-----

This is just a quick introduction, view the [GoDoc][2] for details.

Basic usage example in a web server:
```go
package main

import (
    mjq "github.com/jorbriib/multiple-job-queue"
    "fmt"
    "log"
    "net/http"
    "time"
)
type TestJob struct {
}

func (tj *TestJob) Handle() {
    time.Sleep(2 * time.Second)
    fmt.Println("Job handled")
}

func handler(w http.ResponseWriter, r *http.Request) {
    dispatcher := mjq.GetDispatcher()
    testJob1 := &TestJob{}
    // Dispatch dispatches testJob1 to high queue
    _ = dispatcher.Dispatch(testJob1, "high")
    w.WriteHeader(200)
}

func main() {
    // InitializeQueues creates a default queue with 1 worker
    _ = mjq.InitializeQueues(1,
            mjq.AddQueue("high", 4), // AddQueue creates a high queue with 4 workers
            mjq.AddQueue("medium", 2),
    	    mjq.AddQueue("low", 1))

    http.HandleFunc("/job", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

Basic usage example in a console command:
```go
package main
    
import (
    mjq "github.com/jorbriib/multiple-job-queue"
    "math/rand"
)
    
func main() {
    q := mjq.InitializeQueues(0,
      		mjq.AddQueue("high", 5),
      		mjq.AddQueue("medium", 2),
      		mjq.AddQueue("low", 1))
      
    queues := []string{"high", "medium", "low"}
   
    dispatcher := mjq.GetDispatcher()
    for i := 0; i < 100; i++ {
        randomQueue := rand.Intn(len(queues))
        testJob1 := &TestJob{}
        _ = dispatcher.Dispatch(testJob1, queues[randomQueue])
    
    }
    q.WaitUntilFinish()
}  
    	    

```

[1]: https://blog.golang.org/using-go-modules
[2]: http://godoc.org/github.com/jorbriib/multiple-job-queue