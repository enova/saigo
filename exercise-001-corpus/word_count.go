package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saigo/exercise-001-corpus/corpus"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of input arguments", len(os.Args))
		return
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	results := corpus.Analyze(string(data))

	for _, val := range results {
		fmt.Println(val.Word, val.Count)
	}
}
