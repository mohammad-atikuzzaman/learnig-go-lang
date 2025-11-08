package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5, 6, 7}

// you can run loops file by uncommenting main function

// func main() {

// 	// traditional for loop
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 	}

// 	// While-like loop
// 	fmt.Println("\nWhile-like loop:")
// 	count := 0
// 	for count < 3 {
// 		fmt.Printf("Count: %d\n", count)
// 		count++
// 	}

// 	// Infinite loop with break
// 	fmt.Println("\nInfinite loop with break:")
// 	counter := 0
// 	for {
// 		fmt.Printf("Counter: %d\n", counter)
// 		counter++
// 		if counter >= 3 {
// 			break
// 		}
// 	}

// 	// Loop with continue
// 	fmt.Println("\nLoop with continue (even numbers):")
// 	for i := 0; i < 10; i++ {
// 		if i%2 != 0 {
// 			continue // Skip odd numbers
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Println()

// 	// loop through on array
// 	for index, value := range arr {
// 		fmt.Println(index, value)
// 	}

// 	// Only values
// 	fmt.Println("\nOnly values:")
// 	for _, value := range arr {
// 		fmt.Printf("Value: %d\n", value)
// 	}

// }
