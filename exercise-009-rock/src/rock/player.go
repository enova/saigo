package main

import (
	"math/rand"
)

type Player interface {
	Type() string
	Play() int
}

type Flipper struct {
	MoveA int
	MoveB int
}

func (p Flipper) Type() string {
	return "Flipper"
}

func (p Flipper) Play() int {
    moves := []int{p.MoveA,p.MoveB}
    return moves[rand.Int() % len(moves)]
}

type Obsessed struct {
	Move int
}

func (p Obsessed) Type() string {
	return "Obsessed"
}

func (p Obsessed) Play() int {
	return p.Move
}

type Cyclone struct {
	MoveCount int
}

func (p Cyclone) Type() string{
	return "Cyclone"
}

func (p Cyclone) Play() int {
	p.MoveCount++
	moves := []int{Rock,Paper,Scissors}
	return moves[p.MoveCount % len(moves)]

}



// RandoRex ...
type RandoRex struct {
}

// Type returns the type of the player
func (p RandoRex) Type() string {
	return "RandoRex"
}

// Play returns a move
func (p RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}
