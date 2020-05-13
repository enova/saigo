package corpus

import (
	"regexp"
	"sort"
	"strings"
)

// WordCount stores the word and it's associated count
type WordCount struct {
	Word  string
	Count int
}

// ByCount implements sort.Interface based on the count field.
type ByCount []WordCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func splitWord(word string) []string {
	slice := regexp.MustCompile("[^a-zA-Z0-9']+").Split(word, -1)
	return slice
}

// Analyze takes in a string of words and returns a slice sorted in reverse numerical order based on the word count
func Analyze(data string) []WordCount {
	words := splitWord(string(data))

	wordMap := make(map[string]int)
	wordSlice := make([]WordCount, 0)

	for index, word := range words {
		words[index] = strings.ToLower(word)
		if _, ok := wordMap[words[index]]; ok {
			wordMap[words[index]]++
		} else {
			if words[index] != "" {
				wordMap[words[index]] = 1
			}
		}
	}

	for key, val := range wordMap {
		wc := WordCount{key, val}
		wordSlice = append(wordSlice, wc)
	}

	sort.Sort(sort.Reverse(ByCount(wordSlice)))

	return wordSlice
}
