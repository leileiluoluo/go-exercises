package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		// for each iteration, start a goroutine to print i
		go func() {
			fmt.Println(i)
		}()
	}

	// waiting for 3 seconds
	time.Sleep(time.Second * 3)
}
