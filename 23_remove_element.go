package main

import "fmt"

func removeElement(slice *[]int, element int) {
	*slice = append((*slice)[:element], (*slice)[element+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Before:", slice)
	removeElement(&slice, 4)
	fmt.Println("After:", slice)
}
