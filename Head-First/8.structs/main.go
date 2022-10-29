package main

import "magazine/magazine"

func main() {
	//sub is a struct pointer
	sub := magazine.DefaultSubscriber("ahmed")
	magazine.ApplyDiscount(sub)
	magazine.PrintInfo(sub)
}
