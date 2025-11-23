package main

import (
	"errors"
	"fmt"
	"strconv"
)

// custom error
var NegativeNumberError = errors.New("negative numbers are not allowed")

func squareRoot(n float64) (float64, error) {
	if n <= 0 {
		return 0, NegativeNumberError
	}
	return n * 0.5, nil
}

func parseAge(n string) (int, error) {
	val, err := strconv.Atoi(n)
	if err != nil {
		return val, fmt.Errorf("Failed to purse age", err)
	}
	if val < 0 {
		return 0, NegativeNumberError
	}
	if val > 150 {
		return 0, errors.New("Unrealistic age")
	}
	return val, nil
}

func main() {
	result, err := squareRoot(-2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square root: %.2f\n", result)
	}

	// age, err := parseAge("43")
	ages := []string{"34", "-24", "abc", "343"}

	for _, ageStr := range ages {
		age, err := parseAge(ageStr)
		if err != nil {
			fmt.Printf("Failed to parse '%s': %v\n", ageStr, err)
		} else {
			fmt.Printf("Parsed age: %d\n", age)
		}
	}
}
