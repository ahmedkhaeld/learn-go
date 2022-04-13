package main

import (
	"fmt"
	"sync"
)

// request is type of dummy function, does not receive or return
type request func()

func main() {
	// requests type map because the order in map is random
	requests := map[int]request{}
	for i := 1; i <= 100; i++ {
		// create a function that return a request
		f := func(n int) request {
			return func() {
				fmt.Println("request", n)
			}
		}
		// push the function into the  map
		requests[i] = f(i)
	}
	// as the iteration is done we should have 100 requests stored in the map

	var wg sync.WaitGroup
	max := 10 // batch of 10 requests max
	// loop through the 100 requests
	for i := 0; i < len(requests); i += max {
		// loop through each batch, that include a 10 reqs
		for j := i; j < i+max; j++ {
			wg.Add(1)
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j+1])
		}
		wg.Wait()
		fmt.Println(max, "requests processed")
	}
}
