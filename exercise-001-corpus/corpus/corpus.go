package corpus

import (
	"fmt"
	"sort"
	"strings"
)

// WordCount object that represent the word and its count
type WordCount struct {
	Word  string
	Count int
}

// Analyze returns array of WordCount
func Analyze(s string) []WordCount {
	var wordCount []WordCount

	replacedStr := removeNonAlphaCharacters(s)

	sSlice := strings.Split(replacedStr, " ")

	var sMap = createMap(sSlice)

	for k, v := range sMap {
		wordCount = append(wordCount, WordCount{k, v})
	}

	sort.Slice(wordCount, func(i, j int) bool {
		return wordCount[i].Count > wordCount[j].Count
	})

	printWordCount(wordCount)

	return wordCount
}

// Remove non alphabet characters
func removeNonAlphaCharacters(s string) string {
	r := strings.NewReplacer("\n", "",
		".", "",
		"?", "",
		",", "",
		"!", "",
		"\"", "",
		":", "")

	return r.Replace(s)
}

// Print word count
func printWordCount(wordCount []WordCount) {
	for _, element := range wordCount {
		fmt.Printf("%d %s\n", element.Count, element.Word)
	}
}

// Create map of word and its counter
func createMap(sSlice []string) map[string]int {
	var sMap = make(map[string]int)

	for _, word := range sSlice {
		if word != "" {
			_, ok := sMap[word]
			if ok {
				sMap[word]++
			} else {
				counter := 1
				sMap[word] = counter
			}
		}
	}

	return sMap
}
