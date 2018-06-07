package main

func main() {

	// Create Game
	game := &Game{}

	// Add Players
	game.Add(&RandoRex{})
	game.Add(&PlayerB{})
	game.Add(&PlayerC{})

	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}

	// Display Results
	game.Display()
}
