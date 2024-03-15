package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []int{0, 1, 2}
	for _, i := range values {
		// for each iteration, start a goroutine to print i
		go func(i int) {
			fmt.Println(i)
			done <- true
		}(i)
	}

	// wait for all goroutines to complete
	for i := 0; i < len(values); i++ {
		<-done
	}
}
