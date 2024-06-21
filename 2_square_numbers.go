package main

import (
	"fmt"
	"sync"
)

// Вычисление квадрата числа и отправки в канал
func square(num int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	result := num * num
	resultChan <- result
}

func main() {

	numbers := []int{2, 4, 6, 8, 10}
	resultChan := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, resultChan)
	}

	// Ожидание завершения горутин
	wg.Wait()

	// Закрытие канала
	close(resultChan)

	for result := range resultChan {
		fmt.Println(result)
	}
}
