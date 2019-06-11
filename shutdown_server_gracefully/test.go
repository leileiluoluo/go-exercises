package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var addr = flag.String("server addr", ":8080", "server address")

func main() {
	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "hello")
	})

	// server
	srv := http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	// make sure idle connections returned
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); nil != err {
			log.Fatalf("server shutdown failed, err: %v\n", err)
		}
		log.Println("server gracefully shutdown")

		close(processed)
	}()

	// serve
	err := srv.ListenAndServe()
	if http.ErrServerClosed != err {
		log.Fatalf("server not gracefully shutdown, err :%v\n", err)
	}

	// waiting for goroutine above processed
	<-processed
}
