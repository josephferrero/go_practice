package main

import (
	"fmt"
)

// define the triangle and square structs
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

// define the interface "shape" that is applied
// to types that implement the getArea() func
type shape interface {
	getArea() (float64, string)
}

// main function to instantiate the triangle
// and square struct, and then pipe them into
// the printArea() function
func main() {

	t := triangle{
		height: 10,
		base:   10,
	}

	s := square{
		sideLength: 20,
	}

	printArea(t)
	printArea(s)
}

// this is the printArea func that is available
// to the interface "shape" which is avaible to
// types that implement the getArea() function
func printArea(s shape) {
	area, sh := s.getArea()
	fmt.Println("The area of the", sh, "is", area)

}

// calculates the area of the triangle and returns
// the area as float64, as well as string containing
// the type of shape this is.
func (t triangle) getArea() (float64, string) {
	return (.5 * t.height * t.base), "triangle"
}

// same as above, except for a square.
func (s square) getArea() (float64, string) {
	return (s.sideLength * s.sideLength), "square"
}
