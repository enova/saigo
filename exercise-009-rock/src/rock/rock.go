package main

import (
	"errors"
	"flag"
	"strconv"
	"strings"
)

type Arg struct{
	playerType int
	quantity int
	moves [] int
}

var normalizedPlayerTypesMap = buildNormalizedPlayerTypesMap()
var moveNamesMapToMoves = buildMovesMap()

func buildMovesMap() map[string]int {
	return map[string]int{
		"s": Scissors,
		"p": Paper,
		"r": Rock}
}

func buildNormalizedPlayerTypesMap() map[string]int{
	return map[string]int{
		"randorex": PlayerTypeRandoRex,
		"flipper": PlayerTypeFlipper,
		"obsessed": PlayerTypeObsessed,
		"cyclone": PlayerTypeCyclone}
}

func parseArguments() []Arg {
	playerConfigs := flag.String(
		"players",
		"RandoRex:1,Flipper:1:ps,Obsessed:1:p,Cyclone:1",
		"Specifies players and their types")
	flag.Parse()

	invalidArgument := func(){
		panic(errors.New("invalid moves"))
	}
	var args []Arg
	for _, playerConfig := range strings.Split(*playerConfigs, ",") {
		playerInfo := strings.Split(playerConfig, ":")
		quantity, _ := strconv.Atoi(playerInfo[1])
		playerType := strings.ToLower(playerInfo[0])
		var moves []int
		if len(playerInfo) > 2 {
			for _, m := range strings.Split(playerInfo[2], "") {
				moves = append(moves, moveNamesMapToMoves[strings.ToLower(m)])
			}
		}
		arg := Arg{normalizedPlayerTypesMap[playerType], quantity, moves}
		moveCount := len(arg.moves)
		switch arg.playerType {
		case PlayerTypeRandoRex,PlayerTypeCyclone:
			if moveCount != 0{invalidArgument()}
		case PlayerTypeObsessed:
			if moveCount != 1{invalidArgument()}
		case PlayerTypeFlipper:
			if moveCount != 2{invalidArgument()}
		}
		args = append(args, arg)
	}
	return args
}

func buildPlayers(args []Arg) []*Player {
	var players []*Player
	for _, arg := range args {
		for i := 0; i < arg.quantity; i++ {
			p := buildPlayer(arg.playerType, arg.moves)
			players = append(players, &p)
		}
	}
	return players
}

func simulateGame(players []*Player) {
	var game Game
	for _, p := range players {
		game.Add(p)
	}
	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}
	// Display Results
	game.Display()
}


func main() {
	args := parseArguments()
	players := buildPlayers(args)
	simulateGame(players)
}
