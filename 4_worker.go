package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Worker function to process data from channel
func Worker(ch chan int64, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			data, ok := <-ch
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println(data, "saved")
		case data := <-ch:
			fmt.Println(data)
		}
	}
}

func main() {
	// Parse number of workers from command line
	numWorkers := flag.Int("workers", 4, "Number of workers to start")
	flag.Parse()

	ch := make(chan int64)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	var wg sync.WaitGroup

	// Start the specified number of workers
	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go Worker(ch, ctx, &wg)
	}

	// Main loop to send data to channel
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- time.Now().Unix()
				time.Sleep(time.Second)
			}
		}
	}()

	// Wait for all workers to finish
	wg.Wait()
	cancel()
	log.Println("All workers are done, exiting.")
}
