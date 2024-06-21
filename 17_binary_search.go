package main

import (
	"fmt"
	"os"
)

func checkArray(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			return false
		}
	}
	return true
}
func binarySearch(arr []int, num int) (int, error) {

	if checkArray(arr) == false {
		fmt.Println("Array is not sorted")
		os.Exit(1)
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == num {
			return mid, nil
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, fmt.Errorf("Number not found")
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res, err := binarySearch(arr, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("index in array:", res)
	}

}
