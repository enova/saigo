package main

import (
	"fmt"
	"io/ioutil"
	"./corpus"
)

func main() {
	var str = readFile("7oldsamr.txt")

	corpus.Analyze(str)
}

func readFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	return str
}
