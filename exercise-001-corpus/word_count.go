package main

import (
	"fmt"
	"git.enova.com/knagabhushana/saigo/exercise-001-corpus/corpus"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	fileContent := string(data)
	WordCountPairs := corpus.Analyze(fileContent)

	for _, s := range WordCountPairs {
		fmt.Println(s.Count, " ", s.Word)
	}

}
