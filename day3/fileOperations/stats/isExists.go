package stats

import (
	"fmt"
	"os"
)

func FileIsExist() {
	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Println("File exists!")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist!")
	} else {
		fmt.Println("Error checking file:", err)
	}
}
