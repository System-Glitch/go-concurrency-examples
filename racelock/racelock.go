package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex

	globalVar string = ""
)

func concat(suffix string) {
	mu.Lock()
	globalVar += suffix
	mu.Unlock()
	wg.Done()
}

func main() {

	wg.Add(2)
	go concat("first")
	go concat("second")

	wg.Wait()
	fmt.Println(globalVar)
}
