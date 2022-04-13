package main

import "fmt"

var w = [...]int{1, 2, 3}
var x = []int{0, 0, 0}

func do(a [3]int, b []int) []int {
	a[0] = 4

	b[0] = 3 // []int{3,0,0}

	c := make([]int, 5) // []int{0,0,0,0,0}

	c[4] = 42 // []int{0,0,0,0,42}

	copy(c, b) // copies only 3 elements to c  > {3,0,0,0,42}
	return c
}

func main() {
	y := do(w, x)
	fmt.Println("y", y)
	fmt.Println("w", w)
	fmt.Println("x", x)
}

// output:
//y [3 0 0 0 42]
//w [1 2 3]
//x [3 0 0]
