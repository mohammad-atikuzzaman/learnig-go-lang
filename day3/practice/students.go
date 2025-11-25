package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID     int
	Name   string
	Grades []float64
}


var students = make(map[int]Student)

func AddStudent(id int, name string, grades []float64) {
	students[id] = Student{
		ID:     id,
		Name:   name,
		Grades: grades,
	}
}


func FindStudentByID(id int) (Student, bool) {
	s, exists := students[id]
	return s, exists
}

func CalculateAverageGrade(grades []float64) float64 {
	if len(grades) == 0 {
		return 0
	}

	sum := 0.0
	for _, g := range grades {
		sum += g
	}
	return sum / float64(len(grades))
}

func SaveToCSV(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{"ID", "Name", "Grades"})

	for _, s := range students {
		var gradeStr []string
		for _, g := range s.Grades {
			gradeStr = append(gradeStr, fmt.Sprintf("%.2f", g))
		}
		writer.Write([]string{
			strconv.Itoa(s.ID),
			s.Name,
			strings.Join(gradeStr, "|"),
		})
	}

	return nil
}
// Load students from CSV
func LoadFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Skip header
	for i := 1; i < len(rows); i++ {
		id, _ := strconv.Atoi(rows[i][0])
		name := rows[i][1]

		gradeStr := strings.Split(rows[i][2], "|")
		var grades []float64
		for _, g := range gradeStr {
			val, _ := strconv.ParseFloat(g, 64)
			grades = append(grades, val)
		}

		students[id] = Student{ID: id, Name: name, Grades: grades}
	}

	return nil
}


func main() {

	// Add sample data
	AddStudent(1, "Atik", []float64{80, 90, 85})
	AddStudent(2, "Hasan", []float64{70, 75, 95})

	// Find student
	s, found := FindStudentByID(1)
	if found {
		fmt.Println("Found Student:", s.Name, "Average:", CalculateAverageGrade(s.Grades))
	}

	// Save to CSV
	err := SaveToCSV("students.csv")
	if err != nil {
		fmt.Println("Error saving:", err)
	} else {
		fmt.Println("Saved to students.csv")
	}

	// Load from CSV
	err = LoadFromCSV("students.csv")
	if err != nil {
		fmt.Println("Error loading:", err)
	} else {
		fmt.Println("Loaded from students.csv")
	}

	fmt.Println("All Students:", students)
}
