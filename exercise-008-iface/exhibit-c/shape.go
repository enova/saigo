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

//////////////////////////
// Equilateral Triangle //
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
	return (math.Sqrt(3) / 4) * (t.side * t.side)
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
	// can call same function for two diff types of structs
	s := Square{side: 10.0}
	Efficiency(&s)

	c := Circle{radius: 10.0}
	Efficiency(&c)

	t := Triangle{side: 10.0}
	Efficiency(&t)
}
