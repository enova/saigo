package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// read in file containing json blob
	data, err := ioutil.ReadFile("exhibit-c/chemistry.json")
	if err != nil {
		panic("Error reading file")
	}

	// initialize empty Element struct
	e := Element{}
	// parses the JSON-encoded data and stores the result in the value pointed to by (in this case) &e
	// use struct tags to map json blob to Element struct
	err = json.Unmarshal(data, &e)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name:     %s\n", e.Name)
	fmt.Printf("Symbol:   %s\n", e.Symbol)
	fmt.Printf("Number:   %d\n", e.AtomicNumber)
	fmt.Printf("Weight:   %f\n", e.AtomicWeight)
	fmt.Printf("Category: %s\n", e.Category)
	fmt.Printf("Group:    %d\n", e.Group)
	fmt.Printf("Period:   %d\n", e.Period)
}
