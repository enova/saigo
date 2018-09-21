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

func parseMoves(moves []string) []int{
	var parsedMoves []int
	for _, m := range moves {
		parsedMoves = append(parsedMoves, moveNamesMapToMoves[strings.ToLower(m)])
	}
	return parsedMoves
}

func parseArguments(playerConfigs string) []Arg{
	var args []Arg
	for _, playerConfig := range strings.Split(playerConfigs, ",") {
		playerInfo := strings.Split(playerConfig, ":")
		quantity, _ := strconv.Atoi(playerInfo[1])
		playerType := strings.ToLower(playerInfo[0])
		var moves []int
		if len(playerInfo) > 2 {
			moves = parseMoves(strings.Split(playerInfo[2], ""))
		}
		arg := Arg{normalizedPlayerTypesMap[playerType], quantity, moves}
		args = append(args, arg)
	}
	return args
}

func validateArguments(args []Arg){
	invalidArgument := func(){
		panic(errors.New("invalid moves"))
	}
	for _, arg := range args {
		moveCount := len(arg.moves)
		switch arg.playerType {
		case PlayerTypeRandoRex,PlayerTypeCyclone:
			if moveCount != 0{invalidArgument()}
		case PlayerTypeObsessed:
			if moveCount != 1{invalidArgument()}
		case PlayerTypeFlipper:
			if moveCount != 2{invalidArgument()}

		}
	}
}

func buildGame(players []*Player) *Game{
	var game Game
	for _, p := range players {
		game.Add(p)
	}
	return &game
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

func simulateGame(game *Game) {
	// A Thousand Round-Robins!
	for i := 0; i < 1000; i++ {
		game.RoundRobin()
	}
	// Display Results
	game.Display()
}


func captureArguments() *string {
	playerConfigs := flag.String(
		"players",
		"RandoRex:1,Flipper:1:ps,Obsessed:1:p,Cyclone:1",
		"Specifies players and their types")
	flag.Parse()
	return playerConfigs
}

func main() {
	rawArgs := *captureArguments()
	args := parseArguments(rawArgs)
	validateArguments(args)
	players := buildPlayers(args)
	game := buildGame(players)
	simulateGame(game)
}
