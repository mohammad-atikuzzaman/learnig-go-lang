package operations

import (
	"fmt"
	"os"
)

func AppendNewContentToFile() {
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
}
