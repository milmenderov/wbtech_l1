package main

import "fmt"

// first way swap
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

	// second way swap
	b, c := 8, 12
	fmt.Println("Nums:", b, c)
	b, c = c, b
	fmt.Println("Swap nums:", b, c)
}
