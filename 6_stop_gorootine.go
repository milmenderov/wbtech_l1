package main

import (
	"context"
	"fmt"
	"time"
)

func worker1(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("1 Stopping goroutine...")
			return
		default:
			fmt.Println("1 Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func worker2(ctx context.Context, str string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("2 Stopping goroutine... ", str)
			return
		default:
			fmt.Println("2 Working... ", str)
			time.Sleep(1 * time.Second)
		}
	}
}

var stopFlag bool

func worker3() {
	for {
		if stopFlag {
			fmt.Println("3 Stopping goroutine...")
			return
		}
		fmt.Println("3 Working...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	stopChan := make(chan struct{})
	go worker1(stopChan)
	time.Sleep(5 * time.Second)
	close(stopChan)
	time.Sleep(2 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx, "manual")
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 4*time.Second)
	go worker2(ctxTimeout, "timeout")
	time.Sleep(4 * time.Second)
	defer cancelTimeout()
	time.Sleep(2 * time.Second)

	deadline := time.Now().Add(5 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	go worker2(ctxDeadline, "deadline")
	time.Sleep(4 * time.Second)
	defer cancelDeadline()
	time.Sleep(2 * time.Second)

	go worker3() // not recommended for use
	time.Sleep(5 * time.Second)
	stopFlag = true
	time.Sleep(2 * time.Second)

}
