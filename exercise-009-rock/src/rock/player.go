package main

import (
	"math/rand"
)

// RandoRex information
type RandoRex struct {

}

// Type returns the type of the player
func (p *RandoRex) Type() string {
	return "RandoRex"
}

// Play returns a move
func (p *RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}

// PlayerB information
type PlayerB struct {

}

// Type returns the type of the player
func (p *PlayerB) Type() string {
	return "PlayerB"
}

// Play returns a move
func (p *PlayerB) Play() int {
	choice := rand.Int() % 5
	return choice
}

// PlayerC information
type PlayerC struct {

}

// Type returns the type of the player
func (p *PlayerC) Type() string {
	return "PlayerC"
}

// Play returns a move
func (p *PlayerC) Play() int {
	choice := rand.Int() % 2
	return choice
}

type Player interface {
	Type() string
	Play() int
}