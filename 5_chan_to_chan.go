package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int)

	Seconds := 10
	timeout := time.After(time.Duration(Seconds) * time.Second)
	wg.Add(1)
	go func() {
		i := 0
		for {
			select {
			case <-timeout:
				close(ch)
				wg.Done()
				return
			default:
				ch <- i
				fmt.Println("Send:", i)
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		for val := range ch {
			fmt.Println("Received:", val)
		}
	}()

	wg.Wait()
	fmt.Println("Program finished.")
}
