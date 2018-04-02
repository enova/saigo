package main

import (
	"fmt"
	"github.com/CoderAlison/saigo/exercise-001-corpus/corpus"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("7oldsamr.txt")
	if err != nil {
		return
	}

	dict := corpus.WordCount(string(bs))
	for _, pair := range dict {
		fmt.Println(pair.Freq, " ", pair.Word)
	}
}
