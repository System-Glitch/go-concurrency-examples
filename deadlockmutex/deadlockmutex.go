package main

import (
	"fmt"
	"sync"
)

var (
	mu       sync.Mutex
	wg       sync.WaitGroup
	username string
)

func updateUsername(newName string) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	username = newName
	validateUsername()
	fmt.Println("Name updated to", username)
}

func validateUsername() {
	mu.Lock()
	defer mu.Unlock()
	if username == "Bob" {
		panic("You cannot be called Bob")
	}
}

func main() {
	wg.Add(3)
	for _, v := range []string{"John", "Michael", "Bob"} {
		go updateUsername(v)
	}

	wg.Wait()
}
