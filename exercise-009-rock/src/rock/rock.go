package main

import (
	"flag"

	p "github.com/akash-ksu/saigo/exercise-009-rock/src/rock/players"
	r "github.com/akash-ksu/saigo/exercise-009-rock/src/rock/rules"
)

func main() {

	var cyclonePlayers, flipperPlayers, obsessedPlayers, randorexPlayers *int

	cyclonePlayers = flag.Int("cyclonePlayers", 1, "Number of cyclone players")
	flipperPlayers = flag.Int("flipperPlayers", 1, "Number of cyclone players")
	obsessedPlayers = flag.Int("obsessedPlayers", 1, "Number of cyclone players")
	randorexPlayers = flag.Int("randorexPlayers", 1, "Number of cyclone players")

	flag.Parse()

	// Create Game
	game := &Game{}

	// Add cyclone players
	for i := 0; i < *cyclonePlayers; i++ {
		game.Add(p.CreatePlayer("cyclone"))
	}

	// Add flipper players
	for i := 0; i < *flipperPlayers; i++ {
		game.Add(p.CreatePlayer("flipper", r.Rock, r.Paper))
	}

	for i := 0; i < *obsessedPlayers; i++ {
		game.Add(p.CreatePlayer("obsessed", r.Paper))
	}

	for i := 0; i < *randorexPlayers; i++ {
		game.Add(p.CreatePlayer("randorex"))
	}

	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}

	// Display Results
	game.Display()
}
