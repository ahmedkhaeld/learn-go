package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		wg.Done()
		// reuse wait, increment the list after it should have done waiting
		wg.Add(1)
	}()
	wg.Wait()
}
