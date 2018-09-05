package players

import (
	"fmt"
	"os"
	"strings"
)

type playerBuilder func(parameters ...int) Player

var registeredPlayers = make(map[string]playerBuilder)

func init() {
	registeredPlayers["cyclone"] = cycloneBuilder
	registeredPlayers["flipper"] = flipperBuilder
	registeredPlayers["obsessed"] = obsessedBuilder
	registeredPlayers["randorex"] = randoRexBuilder
}

func cycloneBuilder(parameters ...int) Player {
	if len(parameters) != 0 {
		fmt.Println("Cyclone player should have no moves")
		os.Exit(1)
	}

	return &Cyclone{}
}

func flipperBuilder(parameters ...int) Player {
	if len(parameters) != 2 {
		fmt.Println("Flipper player should have 2 moves")
		os.Exit(1)
	}

	return &Flipper{Move1: parameters[0], Move2: parameters[1]}
}

func obsessedBuilder(parameters ...int) Player {
	if len(parameters) != 1 {
		fmt.Println("Obsessed player should have only 1 move")
		os.Exit(1)
	}

	return &Obsessed{Move: parameters[0]}
}

func randoRexBuilder(parameters ...int) Player {
	if len(parameters) != 0 {
		fmt.Println("RandoRex player should have no moves")
		os.Exit(1)
	}

	return &RandoRex{}
}

func CreatePlayer(player string, parameters ...int) Player {
	player = strings.ToLower(player)
	builderMethod, ok := registeredPlayers[player]
	if !ok {
		fmt.Print("No player of type -", player)
		os.Exit(1)
	}
	return builderMethod(parameters...)
}
