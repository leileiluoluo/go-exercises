package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
}
