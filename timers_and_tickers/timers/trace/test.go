package main

import (
	"os"
	"runtime/debug"
	"time"
)

func main() {
	debug.SetTraceback("system")
	if len(os.Args) <= 1 {
		panic("before")
	}
	for i := 0; i < 10000; i++ {
		time.NewTimer(time.Second)
		// time.AfterFunc(time.Second, func() {})
	}
	panic("after")
}
