package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
*  This program counts the words in a text file and outputs the counts in decreasing order.
*  Command line: word_count <fileName.txt>
 */
func main() {
	fileName := os.Args[1] // extract file name of text

	// get dictionary
	rawWordCount, err := ScanWords(fileName)
	if err != nil {
		panic(err)
	}

	// sort dictionary
	orderedDict := rankByWordCount(rawWordCount)
	for _, pair := range orderedDict {
		fmt.Println(pair.Value, pair.Key)
	}
}

/*
*  ScanWords -- This function takes in a file name and returns a dictionary of word-count pairs
 */
func ScanWords(path string) (map[string]int, error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	dict := make(map[string]int)
	// go through all words, forcing lowercase & removing punctuations
	for scanner.Scan() {
		curWord := strings.ToLower(scanner.Text())      // force lowercase
		curWord = strings.Replace(curWord, ".", "", -1) // remove special characters
		curWord = strings.Replace(curWord, ",", "", -1)
		curWord = strings.Replace(curWord, "?", "", -1)
		curWord = strings.Replace(curWord, ":", "", -1)
		if curWord[0] == '"' {
			curWord = curWord[1:]
		}
		if curWord[len(curWord)-1] == '"' {
			curWord = curWord[:len(curWord)-1]
		}

		if count, isIn := dict[curWord]; isIn {
			dict[curWord] = count + 1 // word has been seen before, increment count
		} else {
			dict[curWord] = 1 // new word, add to dict
		}
	}

	return dict, nil
}

/*
*  Below is an implementation for sorting a map by value
 */
func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Pair struct used for sorting map by value
type Pair struct {
	Key   string
	Value int
}

// PairList struct used to contain Pair types
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
