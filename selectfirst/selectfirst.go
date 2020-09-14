package main

import (
	"context"
	"fmt"
	"time"
)

func message(ctx context.Context, content string, ch chan string, after time.Duration) {
	time.Sleep(after)
	select {
	case <-ctx.Done():
	case ch <- content:
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go message(ctx, "hello", ch1, time.Second*1)
	go message(ctx, "there", ch2, time.Second*2)

	select {
	case <-ctx.Done():
		fmt.Println("Timeout!")
	case str := <-ch1:
		fmt.Println("Received message from ch1:", str)
	case str := <-ch2:
		fmt.Println("Received message from ch2:", str)
	}
	cancel()

	time.Sleep(time.Second * 1)
}
