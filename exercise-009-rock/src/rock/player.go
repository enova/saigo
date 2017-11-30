package main

import (
	"math/rand"
)

var cycloneCount int
// Player ...
type Player interface {
	getType() string
	getPlay() int
}

type Cyclone struct {
}

func (c *Cyclone) getType() string {
	return "Cyclone"
}

func (c *Cyclone) getPlay() int {
	cycloneCount ++
	choice := cycloneCount % 3
	return choice
}

type Obsessed struct {
	move int
}

func (o *Obsessed) getType() string {
	return "Obsessed"
}

func (o *Obsessed) getPlay() int {
	return o.move
}

type Flipper struct {
	move1 int
	move2 int
}

func (f *Flipper) getType() string {
	return "Flipper"
}

func (f *Flipper) getPlay() int {
	choice := rand.Int() % 2
	if choice == 0 {
		return f.move1
	} else {
		return f.move2
	}
}

type RandoRex struct {
}

func (r *RandoRex) getType() string {
	return "RandoRex"
}

func (r *RandoRex) getPlay() int {
	choice := rand.Int() % 3
	return choice
}

// Play returns a move
func Play(p Player) int {
	return p.getPlay()
}

func Type(p Player) string {
	return p.getType()
}
