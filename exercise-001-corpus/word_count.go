package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saigo/exercise-001-corpus/corpus"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(string(filename))
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	results := corpus.Analyze(string(data))

	for _, val := range results {
		fmt.Println(val.Word, val.Count)
	}
}
