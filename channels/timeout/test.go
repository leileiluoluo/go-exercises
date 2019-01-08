package main

import (
	"errors"
	"fmt"
	"time"
)

func sum(nums []int) int {
	rlt := 0
	for _, num := range nums {
		rlt += num
	}
	return rlt
}

func sumWithTimeout(nums []int, timeout time.Duration) (int, error) {
	rlt := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		rlt <- sum(nums)
	}()
	select {
	case v := <-rlt:
		return v, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	timeout := 3 * time.Second // time.Second
	rlt, err := sumWithTimeout(nums, timeout)
	if nil != err {
		fmt.Println("error", err)
		return
	}
	fmt.Println(rlt)
}
