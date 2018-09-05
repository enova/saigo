package main

import "github.com/akash-ksu/saigo/exercise-008-iface/exhibit-c/shapes"

func main() {
	s := shapes.Build("square", 10)
	shapes.Efficiency(s)

	c := shapes.Build("circle", 10)
	shapes.Efficiency(c)

	r := shapes.Build("rectangle", 10.0, 4.0)
	shapes.Efficiency(r)
}
