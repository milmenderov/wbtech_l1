package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	timer := time.NewTimer(t)
	<-timer.C
}

func main() {
	start := time.Now()
	fmt.Println("Program started")

	fmt.Println("Program sleep for 5 seconds")
	Sleep(time.Second * 5)

	fmt.Println("Program ended after:", time.Since(start))
}
