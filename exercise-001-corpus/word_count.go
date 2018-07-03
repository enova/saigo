package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/tharumax/saigo/exercise-001-corpus/corpus"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Filename not specified")
		return
	}

	filename := os.Args[1]

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range corpus.WordCount(string(content)) {
		fmt.Printf("%d %s\n", p.Count, p.Word)
	}
}
