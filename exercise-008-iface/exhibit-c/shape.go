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

//////////////
// Triangle //
//////////////

type Triangle struct {
	side float64
}

func (t *Triangle) Name() string {
	return "Triangle"
}

func (t *Triangle) Perimeter() float64 {
	return t.side * 3
}

func (t *Triangle) Area() float64 {
	return math.Sqrt(3) / 4 * math.Pow(t.side, 2)
}

//////////////
// Octagon //
//////////////

type Octagon struct {
	side float64
}

func (o *Octagon) Name() string {
	return "Octagon"
}

func (o *Octagon) Perimeter() float64 {
	return o.side * 8
}

func (o *Octagon) Area() float64 {
	return 2 * (1 + math.Sqrt(2)) * math.Pow(o.side, 2)
}

// Build a shape of any kind
func Build(shape string, parameters ...float64) Shape {
	switch shape {
	case "Square":
		return &Square{side: parameters[0]}
	case "Circle":
		return &Circle{radius: parameters[0]}
	case "Triangle":
		return &Triangle{side: parameters[0]}
	case "Octagon":
		return &Octagon{side: parameters[0]}
	}
	return nil
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

	t := Triangle{side: 10.0}
	Efficiency(&t)

	o := Octagon{side: 10.0}
	Efficiency(&o)
}
