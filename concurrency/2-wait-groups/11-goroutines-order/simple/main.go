package main

import (
	"fmt"
	"sync"
)

// demo the randomness of order
// the go keyword shows up sequentially
//however when we're on this example
//it's not sequential it's always random
//it's always in its own order
func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("go routine", i+1)
		}(i)
	}
	wg.Wait()
}

//go routine 10
//go routine 5
//go routine 8
//go routine 9
//go routine 1
//go routine 2
//go routine 7
//go routine 6
//go routine 3
//go routine 4
