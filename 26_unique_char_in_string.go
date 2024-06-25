package main

import (
	"fmt"
	"strings"
)

func checkChars(str string) bool {
	char := make(map[rune]bool, 0)
	str = strings.ToLower(str)
	for _, item := range str {
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
	s = "abcdA"
	fmt.Println("string:", s, checkChars(s))
	s = "abcdgf! "
	fmt.Println("string:", s, checkChars(s))

}
