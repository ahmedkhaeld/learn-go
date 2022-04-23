package main

import "fmt"

func main() {
	// items is a slice of arrays of two bytes
	items := [][2]byte{{1, 2}, {2, 2}, {3, 4}}
	var a [][]byte // slice of slice

	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:])
		a = append(a, i)
	}

	fmt.Println(items)
	fmt.Println(a)

	//[[1 2] [2 2] [3 4]]
	//[[1 2] [2 2] [3 4]]

}
