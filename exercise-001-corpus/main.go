package main

import (
	"fmt"
	"github.com/mrhectorcrackle/saigo/exercise-001-corpus/word_count"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must specify a .txt file, eg word_count test.txt")
		return
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Cannot find file:" + os.Args[1])
		return
	}
	str := string(file)

	words := wordCount.Count(str)
	for _, word := range words {
		fmt.Println(word.Count, word.Key)
	}
}
