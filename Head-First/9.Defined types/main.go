package main

import "fmt"

func main() {
	//soda is instance Liters that contains 2 liters
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.toGallons())

	//water is instance of Milliliters that contains 500ml
	water := Milliliters(500)
	fmt.Printf("%0.0f milli-liters equals %0.3f gallons\n", water, water.toGallons())

	//milk is instance of Gallons, that contains 2 gallons of milk
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equal %0.3f liters", milk, milk.toLiters())
	fmt.Printf("%0.3f gallons equal %0.3f milli-liters", milk, milk.toMilliliters())

}

type Milliliters float64

func (ml Milliliters) toGallons() Gallons {
	return Gallons(ml * 0.000264)
}

type Liters float64

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

type Gallons float64

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) toMilliliters() Liters {
	return Liters(g * 3785.41)
}
