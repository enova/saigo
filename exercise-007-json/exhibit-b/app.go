package main

import (
	"encoding/json"
	"fmt"
)

type Element struct {
	// maps the struct field to name field in the json
	Name string `json:"name"`
	// maps the struct field to symbol field in the json
	Symbol string `json:"symbol"`
	// maps the struct field to atomic_number field in the json
	AtomicNumber int `json:"atomic_number"`
	// maps the struct field to automic_weight field in the json
	AtomicWeight float64 `json:"atomic_weight"`
	// maps the struct field to category field in the json
	Category string `json:"category"`
	// maps the struct field to group field in the json
	Group int `json:"group"`
	// maps the struct field to period field in the json
	Period int `json:"period"`
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
