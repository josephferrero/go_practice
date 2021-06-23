package main

import (
	"fmt"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() (float64, string)
}

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

func printArea(s shape) {
	area, sh := s.getArea()
	fmt.Println("The area of the", sh, "is", area)

}

func (t triangle) getArea() (float64, string) {
	return (.5 * t.height * t.base), "triangle"
}

func (s square) getArea() (float64, string) {
	return (s.sideLength * s.sideLength), "square"
}
