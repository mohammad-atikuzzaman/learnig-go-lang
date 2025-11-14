package main

import "fmt"

func main() {
	// map declaration ways

	var map1 map[string]int      //nil
	map2 := make(map[string]int) //empty map
	map3 := map[string]int{      // initialized map
		"apple":  2,
		"banana": 5,
		"mango":  10,
	}
	map4 := map[string]string{
		"name":       "Akash",
		"department": "management",
		"aim":        "software engineer",
	}

	fmt.Println("Map 1:", map1, "Nil?:", map1 == nil)
	fmt.Println("Map 2:", map2)
	fmt.Println("Map 3:", map3)
	fmt.Println(map4)

	// map operations

	// insert/update
	map3["banana"] = 10 //update the number or banana
	map2["lemon"] = 5   // insert lemon and its count

	fmt.Println(map3)
	fmt.Println(map2)

	// access
	fmt.Println(map3["mango"])
	fmt.Println(map2["lemon"])

	// check if key exist
	value, exist:=map3["mango"]
	if exist {
		fmt.Println("The number of mango: ",value)
	}else{
		fmt.Println("Mango not exist")
	}
}
