package main

import (
	"errors"
	"fmt"
)


type DatabaseError struct {
	Operation string
	Message   string
	Code      int
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error [%d] during %s: %s",
		e.Code, e.Operation, e.Message)
}

func connectToDatabase() error {
	return &DatabaseError{
		Operation: "connect",
		Message:   "connection timeout",
		Code:      1001,
	}
}

func processUserData(userID int) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}

	// Simulate database erro
	if userID == 999 {
		return connectToDatabase()
	}

	return nil
}

func main() {
	userIDs := []int{123, -5, 999, 456}

	for _, userID := range userIDs {
		err := processUserData(userID)
		if err != nil {
			fmt.Printf("Error processing user %d: %v\n", userID, err)

			// Type assertion for custom errors
			if dbErr, ok := err.(*DatabaseError); ok {
				fmt.Printf("Database error details - Code: %d, Operation: %s\n",
					dbErr.Code, dbErr.Operation)
			}
		} else {
			fmt.Printf("Successfully processed user %d\n", userID)
		}
	}
}
