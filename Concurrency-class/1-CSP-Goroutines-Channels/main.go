package main

import "fmt"

//generate has a limit, runs until this limit hits, write to channel the numbers
// then read from this channel the numbers to be filtered
func generate(limit int, ch chan<- int) {
	// start of by 2, ignore 1, every time we make a number, put it into the channel
	for i := 2; i < limit; i++ {
		ch <- i
	}
	close(ch) // when get to the end, close the channel
}

//filter each filter needs to know a src chan(a channel to read from, where the numbers
//coming from), dst chan(where write numbers out), prime number to filter on
func filter(src <-chan int, dst chan<- int, prime int) {
	// loop over the src channel value until the channel closes
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst) // when the loop is finished close the destination
}

func sieve(limit int) {
	ch := make(chan int)

	go generate(limit, ch)
	for {
		// read from the channel the generated numbers, loop to get each number,
		// each number to be tested in the filter prime or not.
		prime, ok := <-ch
		if !ok {
			break
		}
		// make a new channel to represent the new filter
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		// pass the new filter to the main
		ch = ch1
		fmt.Print(prime, " ")
	}
}

func main() {
	sieve(100)
	//2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
}
