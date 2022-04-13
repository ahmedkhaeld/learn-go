package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

// implement the batch approach, so that we limit the amount of connections,
// the amount of requests we make to the server
func main() {
	// declare the total and max(batches) number of requests
	total, max := 10, 3
	var wg sync.WaitGroup
	for i := 0; i < total; i += max {
		limit := max
		// batches: [3, 3, 3, 1]
		if i+max > total {
			limit = total - i
			// this to change the limit dynamically to 1 req
			//the remainder of the batches
		}

		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("could not dial: %v", err)
				}

				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("could not read from connection: %v", err)
				}

				if string(bs) != "success" {
					log.Fatalf("request error, request: %d", i+1+j)
				}

				fmt.Printf("request %d: success\n", i+1+j)
			}(j)
		}

		wg.Wait()
	}
}
