# if...else
#### Making the grades
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//getting the grades from the user
	fmt.Print("Enter a grade: ")
	//set up a buffered reader that gets text from keyboard
	reader := bufio.NewReader(os.Stdin)
	//return what  the user has typed as string, up where they pressed enter key
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//input hast two issues: is string with newline character
	//1.strip that newline char from the input
	// use strings.Trimspace which will remove all the whitespace characters(newlines, tabs, regular spaces)
	//2.convert remainder to floating point using strconv.ParseFloat
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if grade >= 60 {
		status := "passing"
		fmt.Print(status)
	} else {
		status := "failing"
		fmt.Print(status)
	}
}
```
#### Game requirements
//1. Generate a random number from 1 to 100
//and store it as a target number for the player to guess

//2.Prompt the player to guess what the target number is,
and store their response

//3. if the player's guess is less than the target number, say "oops ur guess was low"
if the player's guess is greater than the target number, say, "oops your was high"

//4. Allow the player to guess up to 10 times. Before each guess,
let them know how many guesses they have left

//5. if the player's guess is equal to the target number, tell them "Good job"
You guessed it then to asking for new guesses

//6. if the player ran out of turns without guessing correctly, 
say "sorry you did not guess my number, it was [guess]"

```go
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
	t := target()

	success := false
	//make the user try to guess 10 times
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses")
		g := guess()

		// compare current guess to the target
		if g < t {
			fmt.Println("oops. Your guess was low.")
		} else if g > t {
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
	t := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 to 100.")
	fmt.Println("Can You guess it?")
	return t
}

func guess() int {
	fmt.Print("Make a guess:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	g, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return g
}

```