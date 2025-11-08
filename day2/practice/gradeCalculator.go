package main

import "fmt"

type Student struct {
	School string
	Name   string
	Roll   int
	Marks  float64
}

func CalculateGrade(s Student) (float64, string) {
	if s.Marks < 0 || s.Marks > 100 {
		return 0, "Marks will be minimum 0 and maximum 100"
	}

	if s.Marks >= 80 {
		return s.Marks, "A+"
	} else if s.Marks >= 60 {
		return s.Marks, "A"
	} else if s.Marks >= 50 {
		return s.Marks, "B"
	} else if s.Marks >= 40 {
		return s.Marks, "C"
	} else if s.Marks >= 33 {
		return s.Marks, "D"
	} else {
		return 0, "F"
	}

}

func Students() {
	var name string
	var roll int
	var marks float64

	fmt.Println("Please input name: ")
	fmt.Scan(&name)
	fmt.Println("Please input Roll: ")
	fmt.Scan(&roll)
	fmt.Println("Please input Marks: ")
	fmt.Scan(&marks)

	s1 := Student{"Akash codding school", name, roll, marks}
	m, g := CalculateGrade(s1)
	fmt.Printf("Obtaign marks is: %f\n, And Grade is: %s\n", m, g)
}

func main() {
	Students()
}
