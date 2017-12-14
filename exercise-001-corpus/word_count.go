package main

import (
	"os"

	"github.com/enova/saigo/exercise-001-corpus/corpus"
)

func main() {
	fileName := os.Args[1]
	corpus.CountWords(fileName)
}
