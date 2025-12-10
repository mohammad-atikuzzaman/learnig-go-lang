package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Unbuffered channel (synchronous)
	fmt.Println("=== Unbuffered Channel ===")
	ch1 := make(chan string)

	go func() {
		fmt.Println("Sending message...")
		ch1 <- "Hello from goroutine!"
		fmt.Println("Message sent!")
	}()

	time.Sleep(1 * time.Second)
	msg := <-ch1
	fmt.Println("Received:", msg)

	// Buffered channel (asynchronous)
	fmt.Println("\n=== Buffered Channel ===")
	ch2 := make(chan string, 2)

	ch2 <- "First"
	ch2 <- "Second"
	// ch2 <- "Third" // This would block because buffer is full

	fmt.Println("Received:", <-ch2)
	fmt.Println("Received:", <-ch2)

	// Channel directions
	fmt.Println("\n=== Channel Directions ===")
	dataChan := make(chan int)
	resultChan := make(chan int)

	go producer(dataChan)             // Can only send
	go consumer(dataChan, resultChan) // Can receive from dataChan, send to resultChan

	finalResult := <-resultChan
	fmt.Println("Final result:", finalResult)

	// Select statement (like switch for channels)
	fmt.Println("\n=== Select Statement ===")
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "from ch3"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch4 <- "from ch4"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch3:
			fmt.Println("Received:", msg)
		case msg := <-ch4:
			fmt.Println("Received:", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}

	// WaitGroup for synchronization
	fmt.Println("\n=== WaitGroup Example ===")
	var wg sync.WaitGroup
	results := make([]string, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg, &results)
	}

	wg.Wait()
	fmt.Println("All workers completed!")
	fmt.Println("Results:", results)
}

// Send-only channel
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

// Receive from one channel, send to another
func consumer(in <-chan int, out chan<- int) {
	sum := 0
	for num := range in {
		sum += num
		fmt.Println("Processed:", num)
	}
	out <- sum
}

func worker(id int, wg *sync.WaitGroup, results *[]string) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id+1) * time.Second)
	(*results)[id] = fmt.Sprintf("Worker %d completed", id)
	fmt.Printf("Worker %d done\n", id)
}
