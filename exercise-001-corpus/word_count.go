package main

import (
	"./corpus"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need to include a filename!")
		return
	}
	str := fileIntoString(os.Args[1])
	corpus.Analyze(str)
}

func fileIntoString(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
