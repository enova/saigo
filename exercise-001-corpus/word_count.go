package main

import (
	"fmt"
	"github.com/CoderAlison/saigo/exercise-001-corpus/corpus"
	"os"
)

func main() {
	file, err := os.Open("7oldsamr.txt")
	if err != nil {
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	dict := corpus.WordCount(string(bs))
	for _, pair := range dict {
		fmt.Println(pair.Freq, " ", pair.Word)
	}
}
