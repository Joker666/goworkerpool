package main

import (
	"fmt"
	"github.com/Joker666/goworkerpool/basic"
	"github.com/Joker666/goworkerpool/model"
	"github.com/Joker666/goworkerpool/worker"
)

func main() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 200; i++ {
		data := model.SimpleData{ ID: i }
		allData = append(allData, data)
	}
	fmt.Printf("Start processing all work \n")

	// Process
	basic.Work(allData)
	worker.NotPooledWork(allData)
	worker.PooledWork(allData)
	worker.PooledWorkError(allData)
}