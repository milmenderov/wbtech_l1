package main

import (
	"fmt"
	"sync"
)

func squaree(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		sq := num * num
		out <- sq
	}
	close(out)
}

func stdExit(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	in := make(chan int, 5)
	out := make(chan int, 5)
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(2)
	go squaree(in, out, &wg)
	go stdExit(out, &wg)

	for _, val := range numbers {
		in <- val
	}
	close(in)

	wg.Wait()
}
