// Hands-on exercise #5
// ● create a type SQUARE
// ● create a type CIRCLE
// ● attach a method to each that calculates AREA and returns it
// 	○ circle area= π r 2
// 	○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area

package main

import (
	"fmt"
	"math"
)

// -------- Square --------

type square struct {
	side float64
}

func (s square) calcArea() float64 {
	return math.Pow(s.side, 2)
}

// -------- Square --------

// -------- Circle --------

type circle struct {
	radius float64
}

func (c circle) calcArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// -------- Circle --------

// -------- Shape Interface --------

type shape interface {
	calcArea() float64
}

// -------- Shape Interface --------

// -------- Info Function --------

func info(sh shape) {
	var details string

	switch sh.(type) {
	case square:
		details = fmt.Sprintf(
			"\tSide: %v\n\tArea: %v",
			sh.(square).side,
			sh.calcArea(),
		)
	case circle:
		details = fmt.Sprintf(
			"\tRadius: %v\n\tArea: %v",
			sh.(circle).radius,
			sh.calcArea(),
		)
	}

	fmt.Println("This shape has the following details:\n", details)
}

// -------- Info Function --------

func main() {
	shape1 := square{
		side: 5.348,
	}

	info(shape1)

	fmt.Println()

	shape2 := circle{
		radius: 8.4279,
	}

	info(shape2)
}
