package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}
