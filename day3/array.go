package main

import "fmt"

func declarationOfArray() {
	// array declaration ways
	var arrWay1 [3]int
	fmt.Println(arrWay1)

	var arrWay2 = [3]int{1, 2, 3}
	fmt.Println((arrWay2))

	arrWay3 := [3]string{"a", "b", "c"}
	fmt.Println(arrWay3)

	// Compiler counts the elements
	arrWay4 := [...]string{"Akash", "Batas", "Nodi", "Sagor", "Pani"}
	fmt.Println(arrWay4)

	// Array operations
	fmt.Println(len(arrWay1))            //for showing length of array
	fmt.Println(arrWay4[0])              //print first element of array
	fmt.Println(arrWay4[len(arrWay4)-1]) //print last element of array

	// Arrays are value types (copied)
	newArr := arrWay2
	newArr[0] = 5
	fmt.Println(arrWay2)
	fmt.Println(newArr)

	newArr2 := arrWay4
	newArr2[0] = "test"
	fmt.Println(arrWay4)
	fmt.Println(newArr2)
}

func LoopThroughOnArray() {
	arr := [...]string{"Akash", "Batas", "Nodi", "Sagor", "Pani"}

	// Iterating through array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Range loop
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

func main() {
	declarationOfArray()
	LoopThroughOnArray()
}
