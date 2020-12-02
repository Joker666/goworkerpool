package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/Joker666/goworkerpool/model"
)

// PooledWorkError handles tasks while also handling error
func PooledWorkError(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup
	workerPoolSize := 100

	dataCh := make(chan model.SimpleData, workerPoolSize)
	errors := make(chan error, 1000)

	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for data := range dataCh {
				process(data, errors)
			}
		}()
	}

	for i := range allData {
		dataCh <- allData[i]
	}

	close(dataCh)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-errors:
				fmt.Println("finished with error:", err.Error())
			case <-time.After(time.Second * 1):
				fmt.Println("Timeout: errors finished")
				return
			}
		}
	}()

	defer close(errors)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

func process(data model.SimpleData, errors chan<- error) {
	fmt.Printf("Start processing %d\n", data.ID)
	time.Sleep(100 * time.Millisecond)
	if data.ID%29 == 0 {
		errors <- fmt.Errorf("error on job %v", data.ID)
	} else {
		fmt.Printf("Finish processing %d\n", data.ID)
	}
}
