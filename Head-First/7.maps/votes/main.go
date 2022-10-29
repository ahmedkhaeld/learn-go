package main

import (
	"fmt"
	"log"
	"votes/datafile"
)

func main() {
	lines, err := datafile.ReadLines("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println("Candidate     Votes")
	for name, votes := range counts {
		fmt.Printf("%s : %d \n", name, votes)
	}
}
