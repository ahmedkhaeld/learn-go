package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work(wg)
	wg.Wait()
}

// pass by value will panic
func work(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("work is done")
}
