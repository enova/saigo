package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Moves
const (
	Rock = iota
	Paper
	Scissors
)

// Game ...
type Game struct {
	players []Player
	points  []int
}

// Add adds a player to the game
func (g *Game) Add(p Player) {
	g.players = append(g.players, p)
	g.points = append(g.points, 0)
}

// RoundRobin plays a round-robin tournament with all players
func (g *Game) RoundRobin() {

	// Randomize
	rand.Seed(time.Now().UnixNano())

	n := len(g.players)

	// For Each Player
	for i := 0; i < n-1; i++ {

		// For Each Opponent
		for j := i + 1; j < n; j++ {

			a := g.players[i]
			b := g.players[j]

			// Determine Winner
			moveA := a.Play()
			moveB := b.Play()
			winner := Winner(moveA, moveB)

			// Update Scores
			switch {
			case winner == 1:
				g.points[i]++
			case winner == -1:
				g.points[j]++
			}
		}
	}
}

// Display displays the current scores
func (g *Game) Display() {

	// For Each Player
	for i, p := range g.players {
		fmt.Printf("%-15s %5d\n", p.Type(), g.points[i])
	}
}
