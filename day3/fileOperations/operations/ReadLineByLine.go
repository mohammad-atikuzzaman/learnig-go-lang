package operations

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLineByLine(){
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

}