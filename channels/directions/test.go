package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {
	pings, pongs := make(chan string, 1), make(chan string, 1)
	ping(pings, "ping")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
