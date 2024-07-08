package main

import (
	"fmt"
	"reflect"
)

func determineType(v interface{}) string {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	var i = 42
	var s = "hello"
	var b = true
	var c = make(chan int)

	fmt.Println(determineType(i)) // int
	fmt.Println(determineType(s)) // string
	fmt.Println(determineType(b)) // bool
	fmt.Println(determineType(c)) // channel
}
