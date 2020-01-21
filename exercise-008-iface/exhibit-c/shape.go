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

////////////////
//  Rectangle  //
////////////////

type Rectangle struct {
	sideA float64
	sideB float64
}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Perimeter() float64 {
	return r.sideA * 2.0 + r.sideB * 2.0
}

func (r *Rectangle) Area() float64 {
	return r.sideA * r.sideB
}


////////////////
//  Rhombus  //
////////////////

type Rhombus struct {
	side float64
	diagonalA float64
	diagonalB float64
}

func (r *Rhombus) Name() string {
	return "Rhombus"
}

func (r *Rhombus) Perimeter() float64 {
	return r.side * 4.0
}

func (r *Rhombus) Area() float64 {
	return r.diagonalA * r.diagonalB / 2
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

	r := Rectangle{
		sideA: 5,
		sideB: 12,
	}
	Efficiency(&r)

	rh := Rhombus{
		side:      10,
		diagonalA: 30,
		diagonalB: 40,
	}
	Efficiency(&rh)
}
