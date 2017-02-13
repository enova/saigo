package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortByWordCount(wordFreq map[string]int) PairList {
	p := make(PairList, len(wordFreq))
	i := 0
	for k, v := range wordFreq {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}

func main() {

	if (len(os.Args)) != 2 {
		fmt.Println("Please provide one filename as argument.")
		return
	}

	// How can I typecast dat into String without using a new var
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)
	txtData := string(dat)

	// Run the regular expression on string to replace everything but
	// alpha-numericals with space
	var regExp = regexp.MustCompile("[[:^alpha:]]")
	cleanedTxt := regExp.ReplaceAllString(txtData, " ")

	// Build the Map(string, int) for this string
	cleanedTxt = strings.ToLower(cleanedTxt)
	wordArr := strings.Fields(cleanedTxt)
	wordFreq := make(map[string]int)
	for _, v := range wordArr {
		wordFreq[v] = wordFreq[v] + 1
	}

	//Sort the map on Value
	sortedByCount := sortByWordCount(wordFreq)
	for _, pl := range sortedByCount {
		fmt.Println(pl.Value, pl.Key)
	}
}
