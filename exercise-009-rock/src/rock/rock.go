package main

func main() {

	// Create Game
	game := &Game{}

	// Add Players
	game.Add(&RandoRex{})
	game.Add(&RandoRex{})
	game.Add(&Flipper{Rock,Paper})
	game.Add(&Obsessed{Paper})
	game.Add(&Obsessed{Scissors})
	game.Add(&Cyclone{})

	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}

	// Display Results
	game.Display()
}
