package main

import "fmt"

func setString(arr []string) []string {
	result := make([]string, 0)
	set := make(map[string]bool, len(arr))

	for _, val := range arr {
		set[val] = true
	}

	for key := range set {
		result = append(result, key)
	}

	return result
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("Arr: %v\nSet: %v\n", arr, setString(arr))
}
