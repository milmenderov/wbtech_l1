package main

import "fmt"

func checkChars(str string) bool {
	char := make(map[rune]bool, 0)
	for _, item := range str {
		if item == ' ' {
			continue
		}
		_, ok := char[item]
		if ok {
			return false
		}
		char[item] = true
	}
	return true
}

func main() {
	s := "abcd"
	fmt.Println("string:", s, checkChars(s))
	s = "abcda"
	fmt.Println("string:", s, checkChars(s))
	s = "abcd gf"
	fmt.Println("string:", s, checkChars(s))

}
