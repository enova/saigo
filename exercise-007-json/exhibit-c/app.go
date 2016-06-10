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
	data, err := ioutil.ReadFile("exhibit-c/chemistry.json")
	if err != nil {
		panic("Error reading phones.json")
	}

	e := Element{}
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
