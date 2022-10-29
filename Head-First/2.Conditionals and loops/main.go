package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := target()

	success := false
	//make the user try to guess 10 times
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses")
		guess := guess()

		// compare current guess to the target
		if guess < target {
			fmt.Println("oops. Your guess was low.")
		} else if guess > target {
			fmt.Println("oops. Your guess was high")
		} else {

			success = true // to prevent failure message, after going out of the loop
			fmt.Println("Good Job!, You guessed it!")
			// stop asking the user for guesses
			break // break out of the loop when the user get it right
		}
	}
	// when looping is finished, and user consumed all the tries, without getting it right
	if !success {
		fmt.Println("sorry, you did not guess my number. it was:", target)
	}
}

func target() int {
	//we need to pass a value to rand.Seed,
	//so we get different nums
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 to 100.")
	fmt.Println("Can You guess it?")
	return target
}

func guess() int {
	fmt.Print("Make a guess:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return guess
}
