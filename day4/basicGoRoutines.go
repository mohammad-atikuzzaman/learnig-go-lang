package main

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Number: %d\n", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Letter: %c\n", i)
	}
}
func main() {
	fmt.Println("=== Sequential Execution ===")
	start := time.Now()
	PrintNumbers()
	printLetters()
	fmt.Printf("Sequential took: %v\n", time.Since(start))

	fmt.Println("\n=== Concurrent Execution ===")
    start = time.Now()
    
    // Run functions concurrently
    go PrintNumbers()
    go printLetters()
    
    // Wait for goroutines to complete
    time.Sleep(1 * time.Second)
    fmt.Printf("Concurrent took: %v\n", time.Since(start))
}
