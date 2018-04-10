package main

import (
	"github.com/saigo/exercise-008-iface/exhibit-d/shapes"
)

func main() {
	shapes.Efficiency(shapes.Build("Square", 10))
	shapes.Efficiency(shapes.Build("Circle", 10))
	shapes.Efficiency(shapes.Build("Rectangle", 10, 5))
	shapes.Efficiency(shapes.Build("Rectangle", 11, 7))

}
