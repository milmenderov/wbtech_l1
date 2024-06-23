package main

import "fmt"

func intersection(sl1, sl2 []int) []int {
	sets1 := make(map[int]bool, len(sl1))
	sets2 := make(map[int]bool, len(sl2))
	result := make([]int, 0)

	for _, val := range sl1 {
		sets1[val] = true
	}
	for _, val := range sl2 {
		sets2[val] = true
	}
	for key := range sets1 {
		if sets2[key] {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sl2 := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(intersection(sl1, sl2))
}
