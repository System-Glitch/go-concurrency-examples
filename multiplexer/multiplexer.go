package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printHash(data []byte) {
	sum := sha256.Sum256(data)
	hex := hex.EncodeToString(sum[:])
	fmt.Println(hex)
	wg.Done()
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	// Build a big list of random values
	count := 20
	list := make([][]byte, 0, count)
	for i := 0; i < count; i++ {
		bytes := make([]byte, 0, 100)
		for j := 0; j < 100; j++ {
			bytes = append(bytes, byte(rand.Intn(math.MaxUint8)))
		}
		list = append(list, bytes)
	}
	wg.Add(len(list))

	// Multiplex to hash all values in parallel
	for _, v := range list {
		go printHash(v)
	}

	wg.Wait()
}
