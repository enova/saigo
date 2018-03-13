package main

import (
	"fmt"
	"os"

	"github.com/mokpro/saigo/exercise-001-corpus/corpus"
)

func main() {
	text := readFile()
	wordFreqs := corpus.Analyze(text)
	for _, word := range wordFreqs {
		fmt.Println(word.Count, " ", word.Word)
	}
}

func readFile() string {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing!")
		os.Exit(3)
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error in opening file!")
		os.Exit(3)
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error reading file stat!")
		os.Exit(3)
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error reading file!")
		os.Exit(3)
	}

	return string(bs)
}
