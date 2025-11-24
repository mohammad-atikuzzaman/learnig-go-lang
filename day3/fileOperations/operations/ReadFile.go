package operations

import (
	"fmt"
	"os"
)

func ReadFile() {
	// Reading entire file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Unable to read file", err)
		return
	}
	// the data of file is in byte format, so we have convert byte in string
	fmt.Println("This file content is : ", string(data))
}
