package players

// Player ...
type Player interface {
	Type() string
	Play() int
}
