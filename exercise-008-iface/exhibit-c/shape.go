package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
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
	width float64
	height float64
}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}


///////////////
// Triangle //
///////////////

type Triangle struct {
	side1 float64
	side2 float64
	side3 float64
}

func (t *Triangle) Name() string {
	return "Triangle"
}

func (t *Triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.side3
}

func (t *Triangle) Area() float64 {
	p := (t.side1 + t.side2 + t.side3) / 2
	return math.Sqrt(p * (p - t.side1) * (p - t.side2) * (p - t.side3))
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

func EnsureTriangleSidesValid(parameters *[]float64){
	params := append([]float64{}, *parameters...)
	sort.Slice(params, func(i,j int) bool{return params[i] < params[j]})
	if params[0] + params[1] < params[2] {
		panic(errors.New("Invalid triangle sides provided"))
	}
}

func Build(shape string, parameters ...float64) Shape {
	switch(shape){
	case "circle": return &Circle{radius:parameters[0]}
	case "rectangle": return &Rectangle{width: parameters[0], height: parameters[1]}
	case "triangle":
		EnsureTriangleSidesValid(&parameters)
		return &Triangle{side1: parameters[0], side2: parameters[1], side3: parameters[2]}
	case "square": return &Square{side: parameters[0]}
	default:
		panic(errors.New("Invalid shape: " + shape))
	}
}


func main() {
	s := Build("square", 10)
	Efficiency(s)

	c := Build("circle", 1,2,3)
	Efficiency(c)

	r := Build("rectangle", 1,2)
	Efficiency(r)

	t := Build("triangle", 10,9,3)
	Efficiency(t)
}
