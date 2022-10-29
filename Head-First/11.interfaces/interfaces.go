package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Print("Spinning")
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + "coffee pot"
}

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering Up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating up")
}

func test() {
	var device Appliance
	device = Fan("Wind Breeze")
	device.TurnOn()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
}
