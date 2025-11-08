package main

import "fmt"

type Address struct {
	City  string
	Aria  string
	Road  int
	House int
}

type Student struct {
	Name     string
	Age      int
	Address  Address
	IsPassed bool
	GPA      float64
	Roll     int
}

type School struct {
	SchoolName string
	Date       string
	Students   []Student
}

// func main() {
// 	person1 := Student{
// 		Name:     "Akash",
// 		Age:      24,
// 		Address:  Address{"Pabna", "Faridpur", 8, 77},
// 		IsPassed: true,
// 		GPA:      4.44,
// 		Roll:     12,
// 	}
// 	person2 := Student{
// 		Name:     "Atik",
// 		Age:      23,
// 		Address:  Address{"Pabna", "Faridpur", 8, 77},
// 		IsPassed: true,
// 		GPA:      4.45,
// 		Roll:     10,
// 	}
// 	person3 := Student{
// 		Name:     "Bataws",
// 		Age:      24,
// 		Address:  Address{"Pabna", "Faridpur", 8, 77},
// 		IsPassed: false,
// 		GPA:      0,
// 		Roll:     19,
// 	}

// 	school1 := School{
// 		SchoolName: "Sky Codding School",
// 		Date:       "12/12/12",
// 		Students:   []Student{person1, person2, person3},
// 	}

// 	fmt.Print(school1)
// }
