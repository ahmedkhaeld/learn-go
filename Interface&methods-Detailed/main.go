package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func AddToA(a int) func(int) int {
	return func(b int) int {
		return Add(a, b)
	}
}
func main() {

	addTo1 := AddToA(1)(2)

	fmt.Println(addTo1) //3
}
