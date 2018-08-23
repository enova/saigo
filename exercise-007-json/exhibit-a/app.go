package main

import (
	"encoding/json"
	"fmt"
)

type Element struct {
	Name         string
	Symbol       string
	AtomicNumber int
	AtomicWeight float64
	Category     string
	Group        int
	Period       int
}

func main() {

	e := Element{
		Name:         "Gold",
		Symbol:       "Au",
		AtomicNumber: 79,
		AtomicWeight: 196.966,
		Category:     "transition metal",
		Group:        11,
		Period:       6,
	}

	data, err := json.MarshalIndent(e, "<prefix>", "<indent>") // add new params if using marshalindent()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
