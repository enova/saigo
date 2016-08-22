package main

import (
	"math/rand"
)

// Player ...
type Player struct {
}

// Type returns the type of the player
func (p *Player) Type() string {
	return "RandoRex"
}

// Play returns a move
func (p *Player) Play() int {
	choice := rand.Int() % 3
	return choice
}
