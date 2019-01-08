package main

import "fmt"

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	// consumer
	go func() {
		for {
			msg, more := <-messages
			if !more {
				fmt.Println("no more message")
				done <- true
				break
			}
			fmt.Println("message received", msg)
		}
	}()

	// producer
	for i := 0; i < 5; i++ {
		messages <- i
	}
	close(messages)
	<-done
}
