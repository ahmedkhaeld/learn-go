package main

import "fmt"

func negate(myBool *bool) {
	*myBool = !*myBool
}
func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
}
