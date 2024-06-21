package main

import "fmt"

func reverseWords(s string) string {
	var words []string
	var word string

	for _, ch := range s {
		if ch == ' ' {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if len(word) > 0 {
		words = append(words, word)
	}
	for a, b := 0, len(words)-1; a < b; a, b = a+1, b-1 {
		words[a], words[b] = words[b], words[a]
	}
	var result []rune
	for i, word := range words {
		if i > 0 {
			result = append(result, ' ')
		}
		result = append(result, []rune(word)...)
	}

	return string(result)
}

func main() {
	s := "snow dog sun"
	fmt.Println("Before: ", s)
	fmt.Println("After : ", reverseWords(s))

	s = "snow"
	fmt.Println("Before: ", s)
	fmt.Println("After : ", reverseWords(s))

	s = ""
	fmt.Println("Before: ", s)
	fmt.Println("After : ", reverseWords(s))

	s = " snow dog sun «главрыбаǄζ"
	fmt.Println("Before: ", s)
	fmt.Println("After : ", reverseWords(s))
}
