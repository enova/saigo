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

	// creates a json representation of e
	data, err := json.Marshal(&e)
	if err != nil {
		panic(err)
	}

	// creates a json representation of e with no prefix and each line with appropriate indentations of tabs
	indentedData, err := json.MarshalIndent(&e, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	fmt.Println(string(indentedData))
}
