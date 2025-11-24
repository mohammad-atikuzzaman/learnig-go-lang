package operations

import (
	"fmt"
	"os"
)

func WriteFile() {
	// Writing to a file
	content := "Hello, World!\nWelcome to Go programming.\nThis is a test file."
	err := os.WriteFile("test.txt", []byte(content), 0644)

	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
	fmt.Println("Write file successful")

}
