package main

import "fmt"

func main() {
	messages := make(chan string)
	signal := make(chan bool)

	// receive with default
	select {
	case <-messages:
		fmt.Println("message received")
	default:
		fmt.Println("no message received")
	}

	// send with default
	select {
	case messages <- "message":
		fmt.Println("message sent successfully")
	default:
		fmt.Println("message sent failed")
	}

	// muti-way select
	select {
	case <-messages:
		fmt.Println("message received")
	case <-signal:
		fmt.Println("signal received")
	default:
		fmt.Println("no message or signal received")
	}
}
