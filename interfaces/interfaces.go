package interfaces

import (
	"fmt"
	"math"
)

/* Interfaces are named collections of method signatures. */

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
To implement an interface in Go,
we just need to implement all the methods in the interface.
Here we implement geometry on rects and circles.
*/
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
If a variable has an interface type,
then we can call methods that are in the named interface.
Here's a generic measure function talking advantage of this to work on any geometry.
*/
func measure(g geometry) {
	fmt.Println("geometry:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func Measure() {
	fmt.Println()
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	/*
		The circle and rect struct types both implement the geometry interface
		so we can use instances of these structs as arguments to measure.
	*/
	measure(r)
	measure(c)
}
