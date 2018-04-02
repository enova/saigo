package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"./corpus"
	)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("USAGE: word_count <file>")
		return;
	}

	filename := os.Args[1]
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: can't read file", filename)
		return;
	}

	wordFreqList, ok := corpus.WordCounter(string(bs))
	if (ok != nil) {
		fmt.Println("ERROR: wordCounter can't process contents of file", filename)
		return;
	}

	for _, wordFreq := range wordFreqList {
		fmt.Println(wordFreq.Freq, wordFreq.Word)
	}


}