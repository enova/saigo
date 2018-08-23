package main

import (
	"math/rand"
)

// RandoRex ...//
type RandoRex struct {
}

// Type returns the type of the player
func (r *RandoRex) Type() string {
	return "RandoRex"
}

// Play returns a move
func (r *RandoRex) Play() int {
	// maps to const iota values 0..2
	choice := rand.Int() % 3
	return choice
}

// Obsessed ...//
type Obsessed struct {
	pick int
}

// Type returns the type of the player
func (o *Obsessed) Type() string {
	return "Obsessed"
}

// Play returns a move
func (o *Obsessed) Play() int {
	return o.pick
}

// Flipper ...//
type Flipper struct {
	first  int
	second int
}

// Type returns the type of the player
func (f *Flipper) Type() string {
	return "Flipper"
}

// Play returns a move
func (f *Flipper) Play() int {
	// maps to const iota values 0..2
	choice := rand.Int() % 2
	if choice == 1 {
		return f.first
	}
	return f.second
}

// Cyclone ...//
type Cyclone struct {
	pick int
}

// Type returns the type of the player
func (c *Cyclone) Type() string {
	return "Cyclone"
}

// Play returns a move
func (c *Cyclone) Play() int {
	c.pick++
	c.pick = c.pick % 3
	return c.pick
}

// Player ...
type Player interface {
	Type() string
	Play() int
}
