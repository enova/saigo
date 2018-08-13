package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"saigo/exercise-001-corpus/corpus"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		return
	}

	wordArray := corpus.StringToArray(string(data)) // break string into array of words
	counter := corpus.WordCount(wordArray)          // create map of word frequency
	pl := corpus.MapToSortedPL(counter)             // create PairList of map so we can sort by the count

	for _, pair := range pl {
		fmt.Println(pair.Value, pair.Key)
	}
}
