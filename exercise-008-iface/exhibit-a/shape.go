package main

import (
	"fmt"
)

////////////
// Square //
////////////

// define square struct with field side type float64
type Square struct {
	side float64
}

// methods for Square struct
func (s *Square) Name() string {
	return "Square"
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

////////////////
// Efficiency //
////////////////

func Efficiency(s *Square) {
	name := s.Name()      // "Square"
	area := s.Area()      // s.side^2
	rope := s.Perimeter() // 4*s.side

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

func main() {
	s := Square{side: 10.0} // create Square instance
	Efficiency(&s)
}
