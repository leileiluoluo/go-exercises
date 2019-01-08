package main

import (
	"fmt"
	"sync"
	"time"
)

const MaxOutstanding = 2

type Req struct {
	id int
}

func handle(req *Req) {
	time.Sleep(time.Second)
	fmt.Println("handle req", req.id)
}

func Serve(queue chan *Req) {
	var wg sync.WaitGroup
	sem := make(chan int, MaxOutstanding)
	for req := range queue {
		wg.Add(1)
		go func(req *Req) {
			fmt.Println("a goroutine launched")
			defer wg.Done()
			sem <- 1
			handle(req)
			<-sem
		}(req)
	}
	wg.Wait()
}

func ServeWithThroughputLimit(queue chan *Req) {
	var wg sync.WaitGroup
	sem := make(chan int, MaxOutstanding)
	for req := range queue {
		wg.Add(1)
		sem <- 1
		go func(req *Req) {
			fmt.Println("a goroutine launched")
			defer wg.Done()
			handle(req)
			<-sem
		}(req)
	}
	wg.Wait()
}

func main() {
	queue := make(chan *Req, 5)

	// requests
	go func() {
		for i := 0; i < 5; i++ {
			queue <- &Req{i}
		}
		close(queue)
	}()

	// server
	// Serve(queue)
	ServeWithThroughputLimit(queue)
}
