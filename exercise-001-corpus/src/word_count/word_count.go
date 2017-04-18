package main

import (
	"fmt"
	"os"
	"saigo/exercise-001-corpus/src/corpus"
)

func main() {
	argsCLI := os.Args[1:]
	fmt.Println("File path: ", argsCLI)
	for i := 0; i < len(argsCLI); i++ {
		file_name := string(argsCLI[i])
		corpus.WordCount(file_name)
	}
}
