package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var sum float64 = 0
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	count := float64(len(args))
	fmt.Printf("Average: %0.2f \n", sum/count)
}
