package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/stackpath/backend-developer-tests/concurrency"
)

func main() {
	log.SetFlags(log.Ltime)
	ctx := context.Background();
	waitC := make(chan bool)
	go func() {
		for {
			// monitor the number of goroutines
			log.Printf("current routines : %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	// Start Worker Pool.
	advancedPool, err := concurrency.NewAdvancedPool(10, 2)
	if err != nil {
		log.Fatalf("Error creating pool: %v", err)
	}
	type responses struct {
		id    int
		value int
	}

	totalTask := 20
	resultChannel := make(chan responses, totalTask)

	for i := 0; i < totalTask; i++ {
		advancedPool.Submit(ctx, func(ctx context.Context) {
			// simple function to write to stdin and sleep for 5 seconds
			// following which worker will be available for next task
			log.Printf("Running task with id %v", i)
			time.Sleep(5 * time.Second)
			resultChannel <- responses{i, i * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultChannel
		log.Printf("Task %d with results:  %d", res.id, res.value)
	}

	<-waitC
}
