package main

import (
	"fmt"
	"time"
)

func caseTest() {
	newTime := time.Now()
	day := newTime.Weekday()
	hour := newTime.Hour()
	min := newTime.Minute()
	fmt.Println(day, hour, min)

	switch day {
	case time.Saturday:
		fmt.Println("This is weekend")
	case time.Friday:
		fmt.Println("This is Zuma Day")
	case time.Sunday:
		fmt.Println("This is Market day")
	default:
		fmt.Println("This is normal weekday")
	}
	// Switch with conditions
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A+")
	case score >= 80:
		fmt.Println("Grade: A")
	case score >= 70:
		fmt.Println("Grade: B")
	case score >= 60:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// Switch with short statement
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// func main() {
// 	caseTest()
// }
