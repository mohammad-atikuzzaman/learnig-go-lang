package main

import (
	"bufio"
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

	// Reading line by line
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Unable to open file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNumber, line)
		lineNumber++
	}

	// appending to file

	appendText := "\nThis is append text content."
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Unable to open file", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(appendText); err != nil {
		fmt.Println("Error appending to file: ", err)
		return
	}
	fmt.Println("Content append successful")

	// Check if file exists
	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Println("File exists!")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist!")
	} else {
		fmt.Println("Error checking file:", err)
	}
}
