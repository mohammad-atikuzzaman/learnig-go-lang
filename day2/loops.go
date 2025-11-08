package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5, 6, 7}

func main() {

	// traditional for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// loop through on array
	// for index, value := range arr {
	// 	fmt.Println(index, value)
	// }
}
