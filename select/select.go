package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func message(ctx context.Context, ch chan int) {
	defer close(ch)
	for {
		n := rand.Intn(100)
		select {
		case <-ctx.Done():
			return
		case ch <- n:
		}
		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	}
}

func main() {
	ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChannel
		cancel()
	}()

	for i := 0; i < 3; i++ {
		go message(ctx, ch)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("All done!")
			return
		case i, ok := <-ch:
			if ok {
				fmt.Println("Received message:", i)
			}
		}
	}
}
