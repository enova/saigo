package main

import (
        "fmt"
        "math"
)

////////////
// Square //
////////////

// Square struct
type Square struct {
        side float64
}

// Name method for Square
func (s *Square) Name() string {
        return "Square"
}

// Perimeter method for Square
func (s *Square) Perimeter() float64 {
        return 4 * s.side
}

// Area method for Square
func (s *Square) Area() float64 {
        return s.side * s.side
}

////////////
// Circle //
////////////

// Circle struct
type Circle struct {
        radius float64
}

// Name method for Circle
func (c *Circle) Name() string {
        return "Circle"
}

// Perimeter method for Circle
func (c *Circle) Perimeter() float64 {
        return 2 * math.Pi * c.radius
}

// Area method for Circle
func (c *Circle) Area() float64 {
        return math.Pi * c.radius * c.radius
}

////////////
// Triangle //
////////////

// Triangle struct
type Triangle struct {
        side1 float64
        side2 float64
        side3 float64
}

// Name method for Triangle
func (t *Triangle) Name() string {
        return "Triangle"
}

// Perimeter method for Triangle
func (t *Triangle) Perimeter() float64 {
        return t.side1 + t.side2 + t.side3
}

// Area method for Triangle
func (t *Triangle) Area() float64 {
        p := .5 * t.Perimeter()
        return math.Sqrt(p*(p-t.side1)*(p-t.side2)*(p-t.side3))
}

/////////////////////////
// Traditional Hexagon //
/////////////////////////

// Hexagon struct
type Hexagon struct {
        side float64
}

// Name method for Hexagon
func (h *Hexagon) Name() string {
        return "Hexagon"
}

// Perimeter method for Hexagon
func (h *Hexagon) Perimeter() float64 {
        return 6*h.side
}

// Area method for Hexagon
func (h *Hexagon) Area() float64 {
        return (3 * math.Sqrt(3)) / 2 * h.side*h.side
}

////////////////
// Efficiency //
////////////////

// Shape Interface
type Shape interface {
        Name() string
        Perimeter() float64
        Area() float64
}

// Efficiency function
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

        t := Triangle{side1: 3.0, side2: 4.0, side3: 5.0}
        Efficiency(&t)

        h := Hexagon{side: 13.4}
        Efficiency(&h)
}
