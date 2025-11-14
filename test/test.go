package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for i := 0; i <= 100000; i++ {
		sum += i
	}
	elapsed := time.Since(start)
	fmt.Println("Time :", elapsed.Milliseconds(), " ms")
}
