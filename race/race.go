package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup

	globalVar string = ""
)

func concat(suffix string) {
	globalVar += suffix
	wg.Done()
}

func main() {

	wg.Add(2)
	go concat("first")
	go concat("second")

	wg.Wait()
	fmt.Println(globalVar)
}
