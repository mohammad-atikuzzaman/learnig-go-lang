package main

import "fmt"

func processScore(scores []int) (int, float64) {
	if len(scores) == 0 {
		return 0, 0.00
	}
	total := 0
	for _, score := range scores {
		total += score
	}

	average := float64(total) / float64(len(scores))
	return total, average
}
func filterEvenNumber(nums []int) []int {
	var even []int
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	return even
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	total, average := processScore(arr)
	fmt.Println(total, average)

	evens := filterEvenNumber(arr)
	fmt.Println(evens)
}
