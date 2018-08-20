package main

import (
	"fmt"
	"github.com/akash-ksu/saigo/exercise-001-corpus/word_count/corpus"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage -- word_count filename")
		return
	}

	filename := args[0]

	br, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading file", filename)
		return
	}

	contents := string(br)

	pairList := corpus.Analyze(contents)

	printCount(pairList)
}

func printCount(pairList corpus.PairList) {
	for _, pair := range pairList {
		fmt.Println(strconv.Itoa(pair.Value) + " : " + pair.Key)
	}
}
