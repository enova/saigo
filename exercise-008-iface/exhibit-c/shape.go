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

///////////////
// Rectangle //
///////////////

type Rectangle struct {
	side1 float64
	side2 float64

}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.side1 + r.side2)
}

func (r *Rectangle) Area() float64 {
	return r.side1 * r.side2
}

//////////////////////////
// Equalataral triangle //
//////////////////////////

type Triangle struct {
	side float64
}

func (t *Triangle) Name() string {
	return "Triangle"
}

func (t *Triangle) Perimeter() float64 {
	return 3 * t.side
}

func (t *Triangle) Area() float64 {
	return (t.side * t.side) / 4 * math.Sqrt(3)
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

	r := Rectangle{side1: 10.0, side2: 5}
	Efficiency(&r)

	t := Triangle{side: 10.0}
	Efficiency(&t)
}
