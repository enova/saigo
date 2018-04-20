package main

import (
	"fmt"
	"github.com/dwisner-enova/saigo/exercise-001-corpus/corpus"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No file specified")
		return
	}
	filePath := args[0]

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(b)

	for _, s := range corpus.SortedWordCount(&str) {
		fmt.Println(s)
	}
}
