package main

import (
	"fmt"
	"math"
)

func calCulateCircleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Radius cant be negative")
	}

	return math.Pi * radius * radius, nil
}

func calculateTriangleArea(base float64, hight float64) (float64, error) {
	if base <= 0 || hight <= 0 {
		return 0, fmt.Errorf("Base or hight cant be 0 or less then 0")
	}

	return 0.5 * base * hight, nil
}


// func main(){
// 	// area, error := calCulateCircleArea(10)
// 	area, error := calculateTriangleArea(10, 20)
// 	if(error != nil){
// 		fmt.Println(error)
// 	}else{
// 		fmt.Println(area)

// 	}
// }