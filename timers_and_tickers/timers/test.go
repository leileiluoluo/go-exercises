package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("hello")

	timer = time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("world")
	}()
	if timer.Stop() {
		fmt.Println("timer stoped")
	}
}
