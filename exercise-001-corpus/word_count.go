package main

import (
	"flag"
	"fmt"
	"saigo/exercise-001-corpus/corpus"
)

func main() {
	file := flag.String("file", "", "the file to parse")
	flag.Parse()

	wordlist, err := corpus.WordCountFromFileSorted(*file, true)
	if err == nil {
		for _, pair := range wordlist {
			fmt.Printf("%s %d\n", pair.Key, pair.Value)
		}
	}
}
