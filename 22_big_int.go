package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(9099512007891750610)
	b := big.NewInt(9099512008521750610)

	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)

	sum.Add(a, b)
	sub.Sub(a, b)
	mul.Mul(a, b)
	div.Div(a, b)

	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("a + b = %d\n", sum)
	fmt.Printf("a - b = %d\n", sub)
	fmt.Printf("a * b = %d\n", mul)
	fmt.Printf("a / b = %d\n", div)
}
