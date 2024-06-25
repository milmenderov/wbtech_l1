package main

import (
	"fmt"
	"sync"
)

func double(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for num := range in {
		sq := num * 2
		out <- sq
	}
	close(out)
	defer wg.Done()
}

func stdOutput(out <-chan int, wg *sync.WaitGroup) {
	for val := range out {
		fmt.Println(val)
	}
	defer wg.Done()
}

func main() {
	in := make(chan int, 5)
	out := make(chan int, 5)
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(2)
	go double(in, out, &wg)
	go stdOutput(out, &wg)

	for _, val := range numbers {
		in <- val
	}
	close(in)

	wg.Wait()
}
