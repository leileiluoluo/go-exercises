package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	fmt.Println(<-messages, <-messages)
}
