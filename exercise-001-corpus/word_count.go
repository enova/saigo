package main

import (
	"carlos/corpus"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires an argument")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for _, wordCount := range corpus.Analyze(string(file)) {
		fmt.Println(wordCount.Count, wordCount.Word)
	}
}
