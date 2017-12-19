package main

import (
	"os"

	"github.com/wonderw/saigo/exercise-001-corpus/corpus"
)

func main() {
	filepath := os.Args[1]
	corpus.WordCount(filepath)
}
