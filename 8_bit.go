package main

import (
	"fmt"
	"strconv"
)

func WithStrings(b int64, index int) {
	var conv = []rune(strconv.FormatInt(b, 2))
	if len(conv)-1 < index {
		fmt.Println("Index out of bounds")
	} else if conv[len(conv)-1-index] == '0' {
		conv[len(conv)-1-index] = '1'
	} else {
		conv[len(conv)-1-index] = '0'
	}
	result, _ := strconv.ParseInt(string(conv), 2, 64)
	fmt.Println(result)
}

func SetBit(b int64, index int, bit int) int64 {
	if bit == 1 {
		return b | (1 << index)
	}
	return b & ^(1 << index)
}

func SetBitAuto(b int64, index int) int64 {
	var sdvig int64 = 1 << index
	return b ^ (sdvig)
}

func main() {
	var b int64 = 140
	index := 3
	WithStrings(b, index)
	fmt.Println(SetBit(b, index, 0))
	fmt.Println(SetBitAuto(b, index))

}
