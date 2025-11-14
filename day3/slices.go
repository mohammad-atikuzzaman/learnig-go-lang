package main

import "fmt"

func main() {
	var slice1 []int               // nil slice
	slice2 := []int{1, 2, 3, 4, 5} // [1 2 3 4 5]
	slice3 := make([]int, 3)       // [0 0 0] - length 3
	slice4 := make([]int, 3, 5)    // [0 0 0] - length 3, capacity 5

	fmt.Println("Slice 1:", slice1, "Nil?:", slice1 == nil)
	fmt.Println("Slice 2:", slice2, "Length:", len(slice2), "Capacity:", cap(slice2))
	fmt.Println("Slice 3:", slice3, "Length:", len(slice3), "Capacity:", cap(slice3))
	fmt.Println("Slice 4:", slice4, "Length:", len(slice4), "Capacity:", cap(slice4))

	// slice operations
	// append elements
	slice2 = append(slice2, 6)
	fmt.Println("After append: ", slice2)

	// append multiple
	slice2 = append(slice2, 7, 8, 9, 10)
	fmt.Println("After multiple append: ", slice2)

	// slicing operations
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Original:", numbers)
	fmt.Println("numbers[2:5]:", numbers[2:5]) // it will start slicing from index 2 and end at before index 5
	fmt.Println("numbers[:5]:", numbers[:5])   // it will start slicing from index 0 and end at before index 5
	fmt.Println("numbers[5:]:", numbers[5:])   // it will start slicing from index 5 and end after last element
	fmt.Println("numbers[:]:", numbers[:])     // [0 1 2 3 4 5 6 7 8 9]

	// Copy slices
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println("Source:", src)
	fmt.Println("Destination:", dest)

	

	// Slices are reference types
	sliceA := []int{1, 2, 3}
	sliceB := sliceA
	sliceB[0] = 100
	fmt.Println("Slice A:", sliceA) // [100 2 3] - changed!
	fmt.Println("Slice B:", sliceB) // [100 2 3]
}
