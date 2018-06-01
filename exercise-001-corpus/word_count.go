package main

import (
	"./corpus"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	var str = readFile(fileName)

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
