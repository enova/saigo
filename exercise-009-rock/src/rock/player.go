package main

import (
	"math/rand"
)

// Player ...
type Player interface {
	Type() string
	Play() int
}

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

type Flipper struct {
	move1 int
	move2 int
}

func (f *Flipper) Type() string {
	return "Flipper"
}

func (f *Flipper) Play() int {
	moves := []int{f.move1, f.move2}
	return moves[rand.Intn(2)]
}

type Obsessed struct {
	move int
}

func (o *Obsessed) Type() string {
	return "Obsessed"
}

func (o *Obsessed) Play() int {
	return o.move
}

type Cyclone struct {
	count int
}

func (c *Cyclone) Type() string {
	return "RandoRex"
}

func (c *Cyclone) Play() int {
	moves := []int{Rock, Paper, Scissors}
	move := moves[c.count%3]
	c.count++
	return move

}
