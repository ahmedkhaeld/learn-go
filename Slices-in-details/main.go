package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:1]
	//c is a slice of length two capacity two
	//it represents the first two elements of a,
	// and it's got no extra space to put anything
	c := b[0:2:2]

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	//a[0xc0000b8000] = [1 2 3]
	//b[0xc0000b8000] = [1]
	//c[0xc0000b8000] = [1 2]

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	//a[0xc0000b8000] = [1 2 3]
	//c[0xc0000be020] = [1 2 5]
}
