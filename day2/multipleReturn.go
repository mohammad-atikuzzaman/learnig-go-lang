package main

import "fmt"

// multiple return value
func calc(a int, b int) (int, int, int) {
	add := a + b
	sub := a - b
	mul := a * b

	return add, sub, mul
}

// Named return value
func advCalc(x, y int) (sum int, diff int, mul int) {
	sum = x + y // No need to declare - already declared
	diff = x - y
	mul = x * y
	return // Naked return - returns named values
}

// Function returning error handling
func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cant divide by zero")
	}
	return a / b, nil
}


func allF() {
	// sum, sub, mul := calc(2, 3)
	// fmt.Print(sum, sub, mul)

	// sum, sub, mul := advCalc(4, 3)
	// fmt.Print(sum, sub, mul)

	div, error := div(3, 0)
	if error != nil{
		fmt.Println(error)
	}else{
		fmt.Println(div)
	}
}

// func main() {
// 	allF()
// }
