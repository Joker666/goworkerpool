package workerpool

import (
	"fmt"
	"sync"
)

// Worker handles all the work
type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

// NewWorker returns new instance of worker
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
		quit:     make(chan bool),
	}
}

// Start starts the worker
func (wr *Worker) Start(wg *sync.WaitGroup) {
	fmt.Printf("Starting worker %d\n", wr.ID)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range wr.taskChan {
			process(wr.ID, task)
		}
	}()
}

// StartBackground starts the worker in background waiting
func (wr *Worker) StartBackground() {
	fmt.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			process(wr.ID, task)
		case <-wr.quit:
			return
		}
	}
}

// Stop quits the worker
func (wr *Worker) Stop() {
	fmt.Printf("Closing worker %d\n", wr.ID)
	go func() {
		wr.quit <- true
	}()
}
