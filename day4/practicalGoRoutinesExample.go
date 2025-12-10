package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string, ch chan<- string) {
	start := time.Now
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start())
	ch <- fmt.Sprintf("Fetched %s in %v (Status: %d)", url, duration, resp.StatusCode)
}

func processData(id int, data []int, resultChan chan<- map[string]interface{}) {
	fmt.Printf("Worker %d started processing %d items\n", id, len(data))
	time.Sleep(time.Duration(len(data)) * 100 * time.Millisecond)

	sum := 0
	for _, num := range data {
		sum += num
	}

	result := map[string]interface{}{
		"worker_id": id,
		"sum":       sum,
		"count":     len(data),
		"average":   float64(sum) / float64(len(data)),
	}

	resultChan <- result
}

func main() {
	urls := []string{
		"https://example.com",
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://api.github.com",
	}

	urlChan := make(chan string, len(urls))

	for _, url := range urls {
		go fetchURL(url, urlChan)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-urlChan)
	}

	// Example 2: Data processing with workers
	fmt.Println("\n=== Data Processing with Workers ===")

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan map[string]interface{}, 3)

	// Split data and process concurrently
	go processData(1, data[:3], resultChan)  // First 3 items
	go processData(2, data[3:7], resultChan) // Next 4 items
	go processData(3, data[7:], resultChan)  // Last 3 items

	// Collect results
	for i := 0; i < 3; i++ {
		result := <-resultChan
		fmt.Printf("Worker %d: Sum=%d, Count=%d, Avg=%.2f\n",
			result["worker_id"], result["sum"], result["count"], result["average"])
	}
}
