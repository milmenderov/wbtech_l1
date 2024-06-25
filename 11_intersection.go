package main

import "fmt"

func intersection(sl1, sl2 []int) []int {
	sets1 := make(map[int]bool, len(sl1))
	result := make([]int, 0)

	for _, val := range sl1 {
		sets1[val] = true
	}

	for key := range sl2 {
		if sets1[sl2[key]] {
			result = append(result, sl2[key])
		}
	}
	return result
}

func main() {
	sl1 := []int{1, 12, 3, 4, 5, 6, 7, 0}
	sl2 := []int{15, 6, 3, 8, 12, 10}
	fmt.Println(intersection(sl1, sl2))
}
