package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/Joker666/goworkerpool/basic"
	"github.com/Joker666/goworkerpool/model"
)

// PooledWork handles the tasks with worker pool
func PooledWork(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup
	workerPoolSize := 100

	dataCh := make(chan model.SimpleData, workerPoolSize)

	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for data := range dataCh {
				basic.Process(data)
			}
		}()
	}

	for i := range allData {
		dataCh <- allData[i]
	}

	close(dataCh)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}
