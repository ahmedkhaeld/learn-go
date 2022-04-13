package main

import (
	"fmt"
	"time"
)

// create a channel to read from child process

func main() {
	now := time.Now()
	// make a channel
	done := make(chan struct{})

	go func() {
		work()

		// go routine writes to the channel(pipe) work() data
		done <- struct{}{}
	}()

	// join-point, channel sends the work(data), join the main process
	<-done
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
	//printing some stuff
	//elapsed: 501.024235ms
	//done waiting, main exits
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}
