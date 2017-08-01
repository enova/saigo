package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type count struct {
	word  string
	count int
}

type byCount []count

func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].count > a[j].count }

func main() {
	var str string
	var ok bool
	file, err := ioutil.ReadFile("7oldsamr.txt")
	if err != nil {
		str, ok = "", false
	}
	str, ok = string(file), true
	if !ok {
		fmt.Println("File error!")
		return
	}
	var words = buildCountArray(str)
	sort.Sort(byCount(words))
	//fmt.Printf("%+v\n", words)
	for _, elem := range words {
		fmt.Printf("%s\t%d\n", elem.word, elem.count)
	}
}

func buildMap(str string) map[string]int {
	words := make(map[string]int)
	wordsList := strings.Fields(purify(str))
	for _, word := range wordsList {
		words[word]++
	}
	return words
}

func buildCountArray(str string) []count {
	var words = buildMap(str)
	countsList := make([]count, 0)
	for w, c := range words {
		elem := count{word: w, count: c}
		countsList = append(countsList, elem)
	}
	return countsList
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
