package main

import (
	"fmt"
	"runtime"
	"time"
)

// without wait-groups
/*
depends on how much time we install the main function, we could execute the other go routines
*/

func main() {
	fmt.Println("# of cores ", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(time.Second)
	fmt.Println("done waiting")
}

func work(id int) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task", id, "done waiting")
}
