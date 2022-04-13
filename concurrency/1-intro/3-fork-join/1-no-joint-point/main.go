package main

import (
	"fmt"
	"time"
)

// example with missing join point
// will prove adding a no join point actually make
// the main function exit right away

//the main function exit right away, thinking it does not have any more task to execute

func main() {
	go work() // fork point

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done waiting, main exit")
	// missing join point
}

func work() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Print some stuff")
}
