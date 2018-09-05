package players

import "math/rand"

type Flipper struct {
	Move1 int
	Move2 int
}

func (f *Flipper) Type() string {
	return "Flipper"
}

// Play ... Flips a coin to select one of the two moves
func (f *Flipper) Play() int {
	num := rand.Int() % 2
	if num == 0 {
		return f.Move1
	}

	return f.Move2
}
