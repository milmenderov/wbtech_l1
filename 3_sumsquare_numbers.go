package main

import (
	"fmt"
	"sync"
)

// Вычисление квадрата числа и отправка в канал
func Square(num int, wg *sync.WaitGroup, resultChan chan int) {
	result := num * num
	resultChan <- result
	defer wg.Done()
}

func main() {

	numbers := []int{2, 4, 6, 8, 10}
	resultChan := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go Square(num, &wg, resultChan)
	}

	// Ожидание завершения горутин
	wg.Wait()

	// Закрытие канала
	close(resultChan)

	sum := 0
	for result := range resultChan {
		sum += result
	}

	fmt.Println("Sum =", sum)
}
