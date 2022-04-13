package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// with wait-groups

func main() {
	fmt.Println("# of cores ", runtime.NumCPU())
	// declare wg
	var wg sync.WaitGroup
	// 10 go routines added
	wg.Add(10)

	now := time.Now()

	// do work for 10 times
	for i := 0; i < 10; i++ {

		go work(&wg, i+1)
	}

	// wait() as the same layer as add()
	wg.Wait()
	fmt.Println("time elapsed", time.Since(now))
	fmt.Println("done waiting")
}

//work takes in wg as a reference(pointer)
func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task", id, "done waiting")
}
