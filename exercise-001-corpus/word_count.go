package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bguthrie/saigo/exercise-001-corpus/corpus"
)

func checkEquality(ar1, ar2 []string) bool {
	if ar1 == nil || ar2 == nil {
		return false
	}

	if len(ar1) != len(ar2) {
		return false
	}

	for i := range ar1 {
		if ar1[i] != ar2[i] {
			return false
		}
	}

	return true
}

func checkInput(args ...string) (string, error) {
	if !checkEquality(args, []string{"./word_count", "file.txt"}) {
		return "", errors.New(`Please provide two arguments: "./word_count file.txt"`)
	}
	return args[1], nil
}

func main() {
	argsWithProg := os.Args
	fileName, err := checkInput(argsWithProg...)
	if err != nil {
		fmt.Println(err)
	} else {
		output := corpus.ParseFile(fileName)
		fmt.Println(output)
	}
}
