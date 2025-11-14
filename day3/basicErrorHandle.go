package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Custom error
var ErrNegativeNumber = errors.New("negative numbers not allowed")

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeNumber
	}
	// Simple square root calculation
	return n * 0.5, nil
}

func parseAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse age: %w", err)
	}

	if age < 0 {
		return 0, errors.New("age cannot be negative")
	}

	if age > 150 {
		return 0, errors.New("age seems unrealistic")
	}

	return age, nil
}

func main() {
	// Basic error handling
	result, err := squareRoot(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square root: %.2f\n", result)
	}

	// Negative number case
	result, err = squareRoot(-4)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Age parsing examples
	ages := []string{"25", "-5", "abc", "200"}

	for _, ageStr := range ages {
		age, err := parseAge(ageStr)
		if err != nil {
			fmt.Printf("Failed to parse '%s': %v\n", ageStr, err)
		} else {
			fmt.Printf("Parsed age: %d\n", age)
		}
	}
}
