package main
import "fmt"

func main(){
	var c float64

	fmt.Println("Please enter temperature in celsius")
	fmt.Scan(&c)
	f := (c * 9/5) + 32

	fmt.Println("The temperature in fahrenheit is: ", f)
}