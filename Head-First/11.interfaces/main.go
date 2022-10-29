package main

import "fmt"

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("string")
	AcceptAnything(true)

}

func AcceptAnything(any interface{}) {
	fmt.Println(any)

}
