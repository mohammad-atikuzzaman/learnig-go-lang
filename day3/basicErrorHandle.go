package main

import (
	"errors"
	"fmt"
)

// custom error
var NegativeNumberError = errors.New("negative numbers are not allowed")

func squareRoot(n float64) (float64, error) {
	if n <= 0 {
		return 0, NegativeNumberError
	}
	return n * 0.5, nil
}

func main() {
	result, err := squareRoot(-2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square root: %.2f\n", result)
	}
}
