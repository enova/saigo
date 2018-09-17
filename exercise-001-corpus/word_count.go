package main

import (
	"fmt"
	"os"
)

import "./corpus"

func main() {
	elements, err := corpus.Exec(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for _, element := range elements {
		fmt.Printf("%s: %d\n", element.Word, element.Count)
	}
}
