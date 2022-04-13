package main

import (
	"fmt"
	"sync"
	"time"
)

// adding a join point using wait-group
// tell the main function to wait for go routines to execute
// before it actually exits
func main() {
	now := time.Now()
	// declare wait-group
	var wg sync.WaitGroup

	// Add function tells the wait-group how many operations to wait for
	wg.Add(1)

	go func() {
		// Done tells the go routine, I'm done executing work, go ahead exit now
		defer wg.Done()
		work()
	}()

	// wait is the join-point
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}
