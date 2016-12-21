## Description

Remember the game Rock-Paper-Scissors? Well what you have in front of you is a rock-paper-scissors game simulator.
At present there is only one type of player. They all randomly (and uniformly) select their action on each turn.
Your task is to add more types of players to the game.

## Comprehension Task

Take a look at the main application code [src/rock/rock.go](https://github.com/enova/saigo/blob/master/exercise-009-rock/src/rock/rock.go):

```go
package main

func main() {

	// Create Game
	game := &Game{}

	// Add Players
	game.Add(&Player{})
	game.Add(&Player{})
	game.Add(&Player{})

	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}

	// Display Results
	game.Display()
}
```

The steps include:

1. Creating a Game
2. Adding (Three) Players
3. Playing a Round-Robin Tournament (1000 Times)
3. Displaying the Results

Try building and running the app:

```
exercise-009-rock$ go build ./...
exercise-009-rock$ ./rock
```

Read through all the code in [src/rock/](https://github.com/enova/saigo/blob/master/exercise-009-rock/src/rock/)
and explain how this game simulation works to an instructor.

## Engineering Tasks

Your task is to add three new player types to the game:

1. `Obsessed` - This player selects the same move every time
1. `Flipper`  - This player flips a coin to select one of two fixed moves
1. `Cyclone`  - This player cycles through the moves repeatedly

The new `rock.go` should look something like this:

```go
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
```

You must:

1. Change type `Player` to `RandoRex`
1. Create an appropriate `Player` interface
1. Change the signature of `Game.Add()` to accept a `Player`
1. Create the three new player types

## Bonus

Make your application accept command-line arguments that describe
the types of players and the number of each type. Your application
should then instantiate those players (using some sort of factory).

