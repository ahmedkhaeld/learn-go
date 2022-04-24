package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

//ScaleBy scales the line, make it longer
func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)

	return Line{l.Begin, Point{l.End.X, l.End.Y}}
}

func main() {

	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance()) // prints 5

	s2 := side.ScaleBy(2)
	fmt.Println(s2.Distance()) //10

	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(2.5).Distance())
	// 12.5
}
