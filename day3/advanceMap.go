package main

import "fmt"

type Student struct {
	Name string
	Roll int
	GPA  float64
	Pass bool
}

func declarationMaps() {
	s1 := Student{"Akash", 12, 4.44, true}
	// fmt.Println(s1)

	students := make(map[int]Student)
	students[0] = s1
	students[1] = Student{"Batas", 13, 3.33, true}
	students[2] = Student{"Sagor", 14, 3.39, true}
	students[3] = Student{"Nodi", 13, 0.00, false}
	// fmt.Println(students)

	// Access
	value, exist := students[0]
	if exist {
		fmt.Println("The value of key 0: ", value)
	} else {
		fmt.Println("0 not exist")
	}

	// loop through on map
	for key, value := range students {
		fmt.Println(key, value)
	}

	student103 := students[3]
	student103.GPA = 4.44
	student103.Pass = true
	fmt.Println(student103)
}

func nestedMaps() {
	classRoster := map[string]map[string]int{
		"classA": {
			"students": 30,
			"teachers": 2,
		},
		"classB": {
			"students": 25,
			"teachers": 1,
		},
	}
	// fmt.Println(classRoster)
	for className, info := range classRoster {
		fmt.Printf("%s: %d students, %d teachers\n",
			className, info["students"], info["teachers"])
	}
}

func main() {
	// declarationMaps()
	nestedMaps()
}
