package main

import (
	"fmt"
	"time"
)

func do(work int, goroutine int) int {
	time.Sleep(time.Second)
	fmt.Printf("goroutine %d done work %d\n", goroutine, work)
	return work
}

func worker(works <-chan int, results chan<- int, goroutine int) {
	for work := range works {
		results <- do(work, goroutine)
	}
}

func startWorkerPool(works <-chan int, results chan<- int, size int) {
	for i := 0; i < size; i++ {
		go worker(works, results, i)
	}
}

func main() {
	works := make(chan int, 10)
	results := make(chan int, 10)

	startWorkerPool(works, results, 2)
	for i := 0; i < 5; i++ {
		works <- i
	}
	close(works)

	// waiting for results
	for i := 0; i < 5; i++ {
		<-results
	}
}
