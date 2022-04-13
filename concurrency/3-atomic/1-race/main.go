package main

import (
	"fmt"
	"sync"
)

// demo where we actually do not respect this principle of atomicity
//where in the end we actually get a race condition
func main() {
	// to simulate race condition, it is enough, to read/write at the same time
	// in at least two go routines
	var count int32 // shared variable
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		count += 10
	}()
	go func() {
		defer wg.Done()
		count -= 15
	}()
	go func() {
		defer wg.Done()
		count++
	}()
	go func() {
		defer wg.Done()
		count = 0
	}()
	go func() {
		defer wg.Done()
		count = 100
	}()
	wg.Wait()

	fmt.Println("count", count)
	// the count will be non-deterministic when running the main function
	// outputs:
	// count -5
	// count 0
	// count 100
}

// when using race flag to demonstrate it is actually a race condition
// $ go run -race main.go
// WARNING: DATA RACE
