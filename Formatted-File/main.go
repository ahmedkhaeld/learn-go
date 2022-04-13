package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// calculate file size,

	var lc, wc, cc int

	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s))
			cc += len(s)
			lc++

		}
		fmt.Printf("%7d %7d %7d %s \n", lc, wc, cc, fname)

		file.Close()

	}
}
