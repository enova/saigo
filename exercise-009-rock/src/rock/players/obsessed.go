package players

type Obsessed struct {
	Move int
}

func (o *Obsessed) Type() string {
	return "Obsessed"
}

// Play ... Selects the same move every time
func (o *Obsessed) Play() int {
	return o.Move
}
