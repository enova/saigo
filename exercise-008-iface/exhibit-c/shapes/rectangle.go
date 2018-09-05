package shapes

///////////////
// Rectangle //
///////////////
type Rectangle struct {
	Length  float64
	Breadth float64
}

// Implement methods from Shape interface for Rectangle

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}
