package main

import (
	"fmt"
	"os"
)

func main() {
	// Writing to a file
	content := "Hello, World!\nWelcome to Go programming.\nThis is a test file."
	err := os.WriteFile("test.txt", []byte(content), 0644)

	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
	fmt.Println("Write file successful")

	// Reading entire file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Unable to read file", err)
		return
	}
	// the data of file is in byte format, so we have convert byte in string
	fmt.Println("This file content is : ", string(data))
}
