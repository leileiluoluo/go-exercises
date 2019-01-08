package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	done := make(chan bool)
	nums := []int{2, 1, 3, 5, 4}
	go func() {
		time.Sleep(time.Second)
		sort.Ints(nums)
		done <- true
	}()
	<-done
	fmt.Println(nums)
}
