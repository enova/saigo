package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	str, ok := readFromFile("7oldsamr.txt")
	if !ok {
		fmt.Println("File error!")
		return
	}
	words := make(map[string]int)
	words = buildMap(str)
	printInOrder(words)
}

func readFromFile(filename string) (string, bool) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", false
	}
	return string(file), true
}

func buildMap(str string) map[string]int {
	words := make(map[string]int)
	wordsArray := strings.Fields(purify(str))
	for i := 0; i < len(wordsArray); i++ {
		if frequency, ok := words[wordsArray[i]]; ok { //key exists, increment frequency
			words[wordsArray[i]] = frequency + 1
		} else { //key does not exist, add to map with 1 frequency
			words[wordsArray[i]] = 1
		}
	}
	return words
}

func printInOrder(words map[string]int) {
	initialLength := len(words)
	for i := 0; i < initialLength; i++ {
		max := 0
		maxKey := ""
		for key, val := range words {
			if val > max {
				maxKey = key
				max = val
			}
			delete(words, maxKey)
		}
		if max > 0 {
			fmt.Printf("\n%d : %s", max, maxKey)
		}
	}
}

func purify(str string) string {
	str = strings.ToLower(str)

	//remove punctuations
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, ";", "", -1)
	str = strings.Replace(str, ":", "", -1)
	str = strings.Replace(str, "\"", "", -1)
	str = strings.Replace(str, "'", "", -1)
	str = strings.Replace(str, "?", "", -1)
	str = strings.Replace(str, "!", "", -1)
	return str
}
