package players

import (
	r "github.com/akash-ksu/saigo/exercise-009-rock/src/rock/rules"
)

// Cyclone ...
type Cyclone struct {
	move int
}

func (c *Cyclone) Type() string {
	return "Cyclone"
}

// Play ... Cycles through all the moves repeatedly
func (c *Cyclone) Play() int {
	allMoves := []int{r.Rock, r.Paper, r.Scissors}
	moveToPlay := allMoves[c.move]
	c.move++
	if c.move == len(allMoves) {
		c.move = 0
	}
	return moveToPlay
}
