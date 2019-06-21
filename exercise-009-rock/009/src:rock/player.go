package main

import (
	"math/rand"
)


var flag int = 0 //helper var for Cyclone

//Player Interface
type Player interface {
	Play() int
	Name() string
}

// Obsessed player selects the same move every time
type Obsessed struct {
}

// Name implementation for Obsessed player
func (o *Obsessed) Name() string {
	return "Obsessed"
}

// Play implementation for Obsessed player
func (o *Obsessed) Play() int {
	return 1
}

// Flipper player flips a coin to select one of two fixed moves
type Flipper struct {
}

// Name implementation for Flipper player
func (f *Flipper) Name() string {
	return "Flipper"
}

// Play implementation for Flipper player
func (f *Flipper) Play() int {
	choice := rand.Int() % 2
	return choice
}

// Cyclone player cycles through the moves repeatedly
type Cyclone struct {
}

// Name implementation for Cyclone player
func (c *Cyclone) Name() string {
	return "Cyclone"
}

// Play implementation for Cyclone player
func (c *Cyclone) Play() int {
	flag += 1
	if flag > 3 {
		flag = 1
	}
	return flag
}

//Type returns the type of the player
type RandoRex struct {
}

// Name implementation for RandoRex player
func (r *RandoRex) Name() string {
	return "RandoRex"
}

// Play implementation for RandoRex player
func (r *RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}
