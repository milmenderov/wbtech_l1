package main

import "fmt"

func reverseString(s string) string {

	rev := []rune(s)
	for a, b := 0, len(rev)-1; a < b; a, b = a+1, b-1 {
		rev[a], rev[b] = rev[b], rev[a]
	}
	return string(rev)
}
func main() {
	s := "«главрыбаǄζ"
	fmt.Println(reverseString(s))
}
