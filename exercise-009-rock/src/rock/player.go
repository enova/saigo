package main

import (
	"math/rand"
)

// Player interface
type Player interface {
	Type() string
	Play() int
}

// RandoRex player type
type RandoRex struct {
}

// Type method for RandoRex
func (p *RandoRex) Type() string {
	return "RandoRex"
}

// Play method for RandoRex
func (p *RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}

// Obsessed player type
type Obsessed struct {
	move int
}

// Type method for Obsessed
func (p *Obsessed) Type() string {
	return "Obsessed"
}

// Play method for Obsesses
func (p *Obsessed) Play() int {
	return p.move
}

// Flipper player type
type Flipper struct {
	move1 int
	move2 int
}

// Type method for Flipper
func (p *Flipper) Type() string {
	return "Flipper"
}

// Play method for Flipper
func (p *Flipper) Play() int {

	// randomly select one of the two fixed moves
	decision := rand.Int() % 2

	var choice int
	switch decision {
	case 0:
		choice = p.move1
	case 1:
		choice = p.move2
	}
	return choice
}

// Cyclone player type
type Cyclone struct {
	currentMove int // cycles through 0-2
}

// Type method for Cyclone
func (p *Cyclone) Type() string {
	return "Cyclone"
}

// Play method for Cyclone
func (p *Cyclone) Play() int {
	p.currentMove++
	p.currentMove %= 3

	return p.currentMove
}
