## Description
Understand interfaces

## Introduction
Check out this code:

```go
package main

import (
	"fmt"
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

////////////////
// Efficiency //
////////////////

func Efficiency(s *Square) {
	name := s.Name()
	area := s.Area()
	rope := s.Perimeter()

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

func main() {
	s := Square{side: 10.0}
	Efficiency(&s)
}
```

This is a simple console application that computes the _efficiency_ of a
square. If you inspect the computation in the function `Efficiency`, you will
see that the value depends on two quantities: _area_ and _perimeter_.
This code is located in the sub-directory
`exhibit-a/shape.go`. Give it a try:

```
exercise-008-iface$ go run exhibit-a/shape.go
```

From the looks of it, the efficiency of a shape should not really be concerned
with the type of shape (in this case a square). The `Efficiency` function only
needs to know the _name_, _area_ and _perimeter_ of the shape in order to carry
out its task. If a customer of this software suggests this efficiency function
may need to handle other shapes, then we would need a way to make it flexible
enough to accept more than just a square. This is where _interfaces_ come in.

## Comprehension Tasks
Within this exercise there are three directories containing command-line application code:

1. `exhibit-a`
1. `exhibit-b`
1. `exhibit-c`

Explain each exhibit to an instructor.

## Question for C++/Java People!

If you have experience with C++ or Java you may have noticed a huge difference in the way Go
implements interfaces. What is it?

## Engineering Task

Expand on `exhibit-c` by adding another shape or two.

## Bonus

Add a shape-factory that can create a shape given a type and a list of lengths:

```
func Build(shape string, parameters ...float64) Shape
```
