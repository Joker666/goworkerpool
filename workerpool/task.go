package workerpool

import (
	"fmt"
	"sync"
)

// Task encapsulates a work item that should go in a work
type Task struct {
	Err  error
	Data interface{}
	f    func(interface{}) error
}

// NewTask initializes a new task based on a given work function.
func NewTask(f func(interface{}) error, data interface{}) *Task {
	return &Task{f: f, Data: data}
}

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

// Stop quits the worker
func (wr *Worker) Stop() {
	fmt.Printf("Closing worker")
	go func() {
		wr.quit <- true
	}()
}

// Pool is the worker pool
type Pool struct {
	Tasks []*Task

	concurrency int
	collector   chan *Task
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		collector:   make(chan *Task, 1000),
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {
	for i := 1; i <= p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		worker.Start(&p.wg)
	}

	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}
	p.Tasks = []*Task{}
	close(p.collector)

	p.wg.Wait()
}

// AddTask adds task to the pool
func (p *Pool) AddTask(task *Task) {
	p.collector <- task
}

func process(workerID int, task *Task) {
	fmt.Printf("Worker %d processes task %v\n", workerID, task.Data)
	_ = task.f(task.Data)
}
