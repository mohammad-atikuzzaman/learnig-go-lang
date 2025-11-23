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
		Operation: "Connect",
		Message:   "connection timeout",
		Code:      1001,
	}
}

func processUserData(userId int) error {
	if userId < 0 {
		return errors.New("Invalid user id")
	}
	// simulate database error
	if userId >= 0 && userId <= 100 {
		return connectToDatabase()
	}
	return nil
}

func main() {
	userIds := []int{1, 4, -2, 400}
	for _, id := range userIds {
		err := processUserData(id)
		if err != nil {
			fmt.Printf("Error processing user %d: %v\n", id, err)
		}

		// Type assertion for custom errors
		if dbErr, ok := err.(*DatabaseError); ok {
			fmt.Printf("Database error details - Code: %d, Operation: %s\n",
				dbErr.Code, dbErr.Operation)
		} else {
			fmt.Printf("Successfully processed user %d\n", id)
		}
	}
}
