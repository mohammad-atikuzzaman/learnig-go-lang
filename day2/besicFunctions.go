package main

import "fmt"

// learning basic function and multiple returns

// simple function
func greet() {
	fmt.Println("Hello from function greet")
}

// function with parameter

func greetPerson(person string) {
	// fmt.Printf("Hello I am :, %s\n , Type: %T\n", person, person)
	fmt.Println("Hello I am : ", person)
}

// return any value

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}

func all() {
	greet()
	greetPerson("Akash")

	// print returned value
	fmt.Println(add(2, 3))
	fmt.Println(mul(2, 3))

	fmt.Print("hello from all function")
}

// func main() {
// 	all()
// }
