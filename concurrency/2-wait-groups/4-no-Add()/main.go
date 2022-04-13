package main

import (
	"fmt"
	"sync"
	"time"
)

// not call Add() before calling the wait() method
// wait will return immediately
func main() {
	var wg sync.WaitGroup
	// missing call add() method
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("go routine: done")
	}()
	wg.Wait()
	fmt.Println("executed immediately")
}
