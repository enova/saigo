package shapes

////////////
// Square //
////////////

type Square struct {
	Side float64
}

func (s Square) Name() string {
	return "Square"
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}