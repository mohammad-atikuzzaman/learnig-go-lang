package main

import "fmt"

type Student struct {
	Name     string
	Role     int
	Email    string
	GPA      float64
	IsActive bool
}

type ClassRoom struct {
	Students map[int]Student
}

func main() {
	s1 := Student{"Akash", 12, "akash@gmail.com", 4.55, true}

	c1 := ClassRoom{Students: map[int]Student{
		s1.Role: s1,
	}}
	fmt.Println(c1)
}
