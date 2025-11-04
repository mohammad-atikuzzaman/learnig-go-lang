package main

import (
	"fmt"
)

// struct definition (like class in other programming language)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Status    bool
}

// Method on Struct (Receiver function)
func (p Person) GetFullName() string {
    return p.FirstName + " " + p.LastName
}

// Method with pointer receiver (can modify struct)
func (p *Person) UpdateStatus(newStatus bool) {
    p.Status = newStatus
}


func main() {
	person1 := Person{
		FirstName: "Md",
		LastName:  "Akash",
		Age:       12,
		Status:    true,
	}

	fmt.Println(person1)

	person2 := Person{"Atik", "Akash", 12, true}
	person2.UpdateStatus(false)
	fmt.Printf("Person info : %+v\n, Type : %T\n", person2 , person2)


	fmt.Println(person1.GetFullName())
}
