package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Joker666/goworkerpool/basic"
	"github.com/Joker666/goworkerpool/model"
	"github.com/Joker666/goworkerpool/worker"
	"github.com/Joker666/goworkerpool/workerpool"
	"github.com/urfave/cli"
)

func main() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 1000; i++ {
		data := model.SimpleData{ID: i}
		allData = append(allData, data)
	}

	var allTask []*workerpool.Task
	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 5)

	app := &cli.App{
		Name:  "goworkerpool",
		Usage: "check different work loads with worker pool",
		Action: func(c *cli.Context) error {
			fmt.Println("You need more parameters")
			return nil
		},
		Commands: []cli.Command{
			{
				Name:  "basic",
				Usage: "run synchronously",
				Action: func(c *cli.Context) error {
					basic.Work(allData)
					return nil
				},
			},
			{
				Name:  "notpooled",
				Usage: "run without any pooling",
				Action: func(c *cli.Context) error {
					worker.NotPooledWork(allData)
					return nil
				},
			},
			{
				Name:  "pooled",
				Usage: "run with pooling",
				Action: func(c *cli.Context) error {
					worker.PooledWork(allData)
					return nil
				},
			},
			{
				Name:  "poolederror",
				Usage: "run with pooling that handles errors",
				Action: func(c *cli.Context) error {
					worker.PooledWorkError(allData)
					return nil
				},
			},
			{
				Name:  "wpool",
				Usage: "run robust worker pool",
				Action: func(c *cli.Context) error {
					pool.Run()
					return nil
				},
			},
			{
				Name:  "wpoolbg",
				Usage: "run robust worker pool in background",
				Action: func(c *cli.Context) error {
					go func() {
						for {
							taskID := rand.Intn(100) + 20

							if taskID%7 == 0 {
								pool.Stop()
							}

							time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
							task := workerpool.NewTask(func(data interface{}) error {
								taskID := data.(int)
								time.Sleep(100 * time.Millisecond)
								fmt.Printf("Task %d processed\n", taskID)
								return nil
							}, taskID)
							pool.AddTask(task)
						}
					}()
					pool.RunBackground()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
