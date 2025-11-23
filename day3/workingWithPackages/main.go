package main

import (
	"fmt"
	"workingWithPackages/models"
	"workingWithPackages/utils"
)

func main() {
    // Using models package
    user1 := models.NewUser(1, "John Doe", "john@example.com")
    user1.Display()
    
    user2 := models.User{
        Id:    2,
        Name:  "Jane Smith", 
        Email: "jane@example.com",
    }
    user2.Display()
    
    // Using utils package
    sum := utils.Add(10, 5)
    product := utils.Multiply(4, 7)
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
    
    numbers := []float64{85.5, 92.0, 78.5, 96.0}
    average := utils.CalculateAverage(numbers)
    fmt.Printf("Average: %.2f\n", average)
}