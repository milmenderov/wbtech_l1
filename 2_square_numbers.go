package main

import (
	"fmt"
	"sync"
)

// Вычисление квадрата числа
func square(num int, wg *sync.WaitGroup) {
	result := num * num
	fmt.Println(result)
	defer wg.Done()

}

func main() {

	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg)
	}

	// Ожидание завершения горутин
	wg.Wait()
}
