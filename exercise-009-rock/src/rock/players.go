package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)
func fail() {
	fmt.Println("Add players using playType,count,picks)\ne.g. ./rock RandoRex,2 Flipper,1,Rock,Paper Obsessed,3,Paper")
	os.Exit(1)
}

func mapMove(move string) int {
	switch move {
	case "Rock": return 0
	case "Paper": return 1
	case "Scissors": return 2
	default: fail()
	}
	return 0
}

func count(expectedLen int, pick []string) int {
	if len(pick) != expectedLen {
		fail()
	}
	count, err := strconv.Atoi(pick[1])
	if err != nil {
		fail()
	}
	return count
}

func (g *Game) addRandoRex(pick []string) {
	count := count(2, pick)
	for i:=1; i <= count; i++ {
		g.Add(&RandoRex{})
	}
}

func (g *Game) addFlipper(pick []string) {
	count := count(4, pick)
	for i:=1; i <= count; i++ {
		pickOne := mapMove(pick[2])
		pickTwo := mapMove(pick[3])
		g.Add(&Flipper{pickOne,pickTwo})
	}
}

func (g *Game) addObsessed(pick []string) {
	count := count(3, pick)
	for i:=1; i <= count; i++ {
		pick := mapMove(pick[2])
		g.Add(&Obsessed{pick})
	}
}

func (g *Game) addCyclone(pick []string) {
	count := count(2, pick)
	for i:=1; i <= count; i++ {
		g.Add(&Cyclone{})
	}
}

func (g *Game) AddPlayers() {
  picks := os.Args
  if len(picks) == 1 {
    fail()
  }
	for i := 1; i < len(picks); i++ {
	  pick := strings.Split(picks[i], ",")
	  switch pick[0] {
	  case "RandoRex": g.addRandoRex(pick)
	  case "Flipper": g.addFlipper(pick)
	  case "Obsessed": g.addObsessed(pick)
	  case "Cyclone":  g.addCyclone(pick)
	  default: fail()
	  }
	}
}
