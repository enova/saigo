package main

func main() {
	// Create Game
	game := &Game{}

	// Add Players
	game.AddPlayers()

	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}

	// Display Results
	game.Display()
}
