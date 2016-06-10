package main

import (
	"encoding/json"
	"fmt"
)

type Element struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	AtomicNumber int     `json:"atomic_number"`
	AtomicWeight float64 `json:"atomic_weight"`
	Category     string  `json:"category"`
	Group        int     `json:"group"`
	Period       int     `json:"period"`
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

	data, err := json.Marshal(&e)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
