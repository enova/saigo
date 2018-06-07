package main

import (
	"math/rand"
)

// RandoRex information
type RandoRex struct {

}

// Type returns the type of the player
func (r *RandoRex) Type() string {
	return "RandoRex"
}

// Play returns a move
func (r *RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}

// Flipper information
type Flipper struct {
	choice1 int
	choice2 int
}

var flipperLastChoice int

// Type returns the type of the player
func (f *Flipper) Type() string {
	return "Flipper"
}

// Play returns a move
func (f *Flipper) Play() int {
	if flipperLastChoice != f.choice1 {
		flipperLastChoice = f.choice1
		return f.choice1
	} else {
		flipperLastChoice = f.choice2
		return f.choice2
	}
}

// Obsessed information
type Obsessed struct {
	choice int
}

// Type returns the type of the player
func (o *Obsessed) Type() string {
	return "Obsessed"
}

// Play returns a move
func (o *Obsessed) Play() int {
	return o.choice
}

var cycloneLastMove = 1

// Cyclone information
type Cyclone struct {
}

// Type returns the type of the player
func (c *Cyclone) Type() string {
	return "Cyclone"
}

// Play returns a move
func (c *Cyclone) Play() int {
	choice := cycloneLastMove
	cycloneLastMove++
	if cycloneLastMove == 3 {
		cycloneLastMove = 1
	}

	return choice
}

type Player interface {
	Type() string
	Play() int
}