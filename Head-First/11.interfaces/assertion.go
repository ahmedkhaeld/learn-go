package main

import "fmt"

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep peep")
}
func (r Robot) Walk() {
	fmt.Println("walking...")
}

type NoiseMaker interface {
	MakeSound()
}

func assert() {
	//define a var with an interface type
	//with assigned value of a type that satisfies the interface
	var noiseMaker NoiseMaker = Robot("robotics")
	noiseMaker.MakeSound() //call the method that's part of the interface

	//convert back to the concrete type using a type assertion
	var robot = noiseMaker.(Robot)
	//call a method that's defined on the concrete type(not the interface)
	robot.Walk()
}
