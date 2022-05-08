package main

import (
	"log"
	"time"
)

// create two things:
//1. stopper, which is worth of five ticks(time out), gives a signal when time is up
//2. ticker, which is a periodic timer
func main() {
	log.Println("start")

	const tickRate = 2 * time.Second

	stopper := time.After(5 * tickRate)  //stopper is a writer channel
	ticker := time.NewTicker(tickRate).C // ticker is a writer channel

loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}

	log.Println("finished")

}
