package main

import "fmt"

func setString(arr []string) []string {
	res := make([]string, 0)
	setOfString := make(map[string]bool, len(arr))

	for _, val := range arr {
		setOfString[val] = true
	}

	for key := range setOfString {
		res = append(res, key)
	}

	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("Arr: %v\nSet: %v\n", arr, setString(arr))
}
