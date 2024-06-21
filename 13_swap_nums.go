package main

import "fmt"

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
func main() {
	a, b := 9223372036854775807, -9223372036854775807
	fmt.Println("Nums :", a, b)
	swap(&a, &b)
	fmt.Println("Swap nums :", a, b)
}
