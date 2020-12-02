package main

import (
	"fmt"
	"github.com/Joker666/goworkerpool/basic"
	"github.com/Joker666/goworkerpool/model"
	"github.com/Joker666/goworkerpool/worker"
	"github.com/Joker666/goworkerpool/workerpool"
	"time"
)

func main() {
	robust()
}

func nonRobust() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 100; i++ {
		data := model.SimpleData{ID: i}
		allData = append(allData, data)
	}
	fmt.Printf("Start processing all work \n")

	// Process
	basic.Work(allData)
	worker.NotPooledWork(allData)
	worker.PooledWork(allData)
	worker.PooledWorkError(allData)
}

func robust() {
	var allTask []*workerpool.Task
	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			fmt.Printf("Task %d processed\n", taskID)
			time.Sleep(100 * time.Millisecond)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 10)
	pool.Run()
}