package basic

import (
	"fmt"
	"time"

	"github.com/Joker666/goworkerpool/model"
)

// Work does the heavy lifting
func Work(allData []model.SimpleData) {
	start := time.Now()
	for i := range allData {
		Process(allData[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

// Process handles the job
func Process(data model.SimpleData) {
	fmt.Printf("Start processing %d\n", data.ID)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finish processing %d\n", data.ID)
}
