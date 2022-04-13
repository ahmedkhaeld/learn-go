package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// do we have some args, and if don't stop
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}
	//remember the first argument is
	//always the name of the program

	// so looking for the second and third args
	old, new := os.Args[1], os.Args[2]

	// scan the input from the cmd
	scan := bufio.NewScanner(os.Stdin)

	// while true, omit the old of the scanned text
	// then replace with new back to the text
	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}

// $ go run . matt hamo < text.txt
// this will replace all the "matt" to "hamo
