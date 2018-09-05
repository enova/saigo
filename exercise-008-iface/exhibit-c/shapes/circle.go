package shapes

import "math"

////////////
// Circle //
////////////

type Circle struct {
	Radius float64
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
