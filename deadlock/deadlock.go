package main

import (
	"fmt"
	"strconv"
)

func toStr(c chan string, val int) {
	if val > 0 {
		c <- strconv.Itoa(val)
	}
}

func main() {
	str := ""

	c1 := make(chan string)
	c2 := make(chan string)
	go toStr(c1, 5)
	go toStr(c2, -1)

	str += <-c1
	str += <-c2
	fmt.Println(str)
}
