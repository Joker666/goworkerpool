package workerpool

import (
	"fmt"
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

func process(workerID int, task *Task) {
	fmt.Printf("Worker %d processes task %v\n", workerID, task.Data)
	_ = task.f(task.Data)
}
