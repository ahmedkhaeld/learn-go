// package main implements wait-group
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type waitGroup struct {
	counter int64
}

// Add increment wg.counter
func (wg *waitGroup) Add(n int64) {
	atomic.AddInt64(&wg.counter, n)
}

// Done decrement wg.counter
func (wg *waitGroup) Done() {
	wg.Add(-1)
	if atomic.LoadInt64(&wg.counter) < 0 {
		panic("negative wait group counter")
	}
}

// Wait exit when wg.counter=0
func (wg *waitGroup) Wait() {
	for {
		if atomic.LoadInt64(&wg.counter) == 0 {
			return
		}
	}
}

func main() {
	var wg waitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("go routine 1")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		fmt.Println("go routine 2")
	}()
	wg.Wait()
	fmt.Println("all go routines are done")
}
