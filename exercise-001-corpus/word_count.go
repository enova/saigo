package main

import (
	"fmt"
	"os"

	"github.com/bguthrie/saigo/exercise-001-corpus/corpus"
)

func main() {
	args := os.Args
	necessaryInputs := []string{"./word_count", "file.txt"}
	errorMessage := "Please input: ./word_count file.txt"

	if args == nil || necessaryInputs == nil {
		panic(errorMessage)
	}

	if len(args) != len(necessaryInputs) {
		panic(errorMessage)
	}

	for i := range args {
		if args[i] != necessaryInputs[i] {
			panic(errorMessage)
		}
	}

	output, err := corpus.ParseFile(args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}
