package main

import (
	"errors"
	"math/rand"
	"reflect"
)


const (
	PlayerTypeRandoRex int = iota
	PlayerTypeFlipper
	PlayerTypeObsessed
	PlayerTypeCyclone
)

// Player ...
type Player interface {
	Type() string
	Play() int
}

type RandoRex struct {}

type Flipper struct {
	moveA int
	moveB int
}
type Obsessed struct {
	move int
}
type Cyclone struct {
	moveState int
}


func (p *RandoRex) Type() string { return getType(p) }
func (p *Flipper) Type() string { return getType(p) }
func (p *Obsessed) Type() string { return getType(p) }
func (p *Cyclone) Type() string { return getType(p) }

func getType(p interface{}) string {
	return reflect.ValueOf(p).Type().Elem().Name()
}

func (p *RandoRex) Play() int {
	return []int{Rock, Paper, Scissors}[rand.Int() % 2]
}

func (p *Flipper) Play() int {
	return []int{p.moveA, p.moveB}[rand.Int() % 2]
}
func (p *Obsessed) Play() int {
	return p.move
}

func (p *Cyclone) Play() int {
	p.moveState++
	return []int{Rock, Paper, Scissors}[p.moveState % 3]
}


func buildPlayer(playerType int, moves []int) Player {
	switch playerType {
	case PlayerTypeRandoRex:
		return &RandoRex{}
	case PlayerTypeFlipper:
		return &Flipper{moves[0], moves[1]}
	case PlayerTypeObsessed:
		return &Obsessed{moves[0]}
	case PlayerTypeCyclone:
		return &Cyclone{}
	default:
		panic(errors.New("invalid type provided"))
	}
}

