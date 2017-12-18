package corpus

import (
	"fmt"
	"sort"
	"strings"
)

// WordCount represents a count of a string/word
type WordCount struct {
	Word  string
	Count int
}

// Analyze will determine a word count of a given string
func Analyze(s string) []WordCount {
	var m []WordCount
	for _, word := range strings.Fields(s) {
		word = replace(word)
		var indexOfWord = contains(m, word)
		if indexOfWord != -1 {
			m[indexOfWord] = WordCount{word, m[indexOfWord].Count + 1}
		} else {
			m = append(m, WordCount{word, 1})
		}
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].Count > m[j].Count
	})
	printWC(m)
	return m
}

func replace(word string) string {
	var replacer = strings.NewReplacer("!", "", ",", "", ".", "", "?", "")
	word = replacer.Replace(word)
	return word
}

func contains(words []WordCount, word string) int {
	index := -1
	for i := range words {
		if words[i].Word == word {
			index = i
		}
	}
	return index
}

func printWC(w []WordCount) {
	fmt.Printf("\n")
	for k := 0; k < len(w); k++ {
		fmt.Printf("%#v   %#v\n", w[k].Count, w[k].Word)
	}
	fmt.Printf("\n")
}
