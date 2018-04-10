package shapes

///////////////
// Rectangle //
///////////////

type Rectangle struct {
	Length float64
	Width float64
}

func (s Rectangle) Name() string {
	return "Rectangle"
}

func (s Rectangle) Perimeter() float64 {
	return s.Length + s.Width
}

func (s Rectangle) Area() float64 {
	return s.Length * s.Width
}
