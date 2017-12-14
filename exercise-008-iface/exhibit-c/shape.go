package main

import (
	"fmt"
	"math"
)

////////////
// Square //
////////////

type Square struct {
	side float64
}

func (s *Square) Name() string {
	return "Square"
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

////////////
// Circle //
////////////

type Circle struct {
	radius float64
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

/////////////
// Ellipse //
/////////////

type Ellipse struct {
	major float64
	minor float64
}

func (e *Ellipse) Name() string {
	return "Ellipse"
}

func (e *Ellipse) Perimeter() float64 {
	// Rough approximation
	return 2 * math.Pi * math.Sqrt((math.Pow(e.major, 2)+math.Pow(e.minor, 2))/2)
}

func (e *Ellipse) Area() float64 {
	return math.Pi * e.major * e.minor
}

////////////////
// Efficiency //
////////////////

type Shape interface {
	Name() string
	Perimeter() float64
	Area() float64
}

func Efficiency(s Shape) {
	name := s.Name()
	area := s.Area()
	rope := s.Perimeter()

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

func main() {
	s := Square{side: 10.0}
	Efficiency(&s)

	c := Circle{radius: 10.0}
	Efficiency(&c)

	e := Ellipse{major: 100, minor: 70}
	Efficiency(&e)
}
