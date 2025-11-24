package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade string
}

func WriteStudentToCSV(students []Student, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Name", "Age", "Grade"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data
	for _, student := range students {
		record := []string{
			student.Name,
			strconv.Itoa(student.Age),
			student.Grade,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func ReadStudentsFromCSV(fileName string) ([]Student, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	
	var students []Student

	// Skip header (index 0)
	for i := 1; i < len(records); i++ {
		age, _ := strconv.Atoi(records[i][1])
		student := Student{
			Name:  records[i][0],
			Age:   age,
			Grade: records[i][2],
		}
		students = append(students, student)
	}

	return students, nil
}

func main() {
	students := []Student{
		{"John Doe", 20, "A"},
		{"Jane Smith", 22, "B"},
		{"Bob Brown", 19, "A+"},
	}

	// Write to CSV
	err := WriteStudentToCSV(students, "students.csv")
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}
	fmt.Println("CSV file created successfully!")

	// Read from CSV
	loadedStudents, err := ReadStudentsFromCSV("students.csv")
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	fmt.Println("\nLoaded students:")
	for _, student := range loadedStudents {
		fmt.Printf("Name: %s, Age: %d, Grade: %s\n",
			student.Name, student.Age, student.Grade)
	}
}
