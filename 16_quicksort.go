package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	mid := len(arr) / 2

	arr[mid], arr[right] = arr[right], arr[mid]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{110, 7, 86, 9, 1, 5, 123, 12, 2}
	fmt.Println("Array:", arr)
	fmt.Println("Sorted array:", quickSort(arr))
}
