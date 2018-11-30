// Package corpus provides basic fucntions for sortin.
package corpus

import (
	"bufio"
	"regexp"
	"sort"
	"strings"
)

// Pair ...(Represents a word as Key and count as Value)
type Pair struct {
	Word  string
	Count int
}

// WordCountPairs ...(custom type that stores a slice of Word Count pairs : (word count))
type WordCountPairs []Pair

func (wc WordCountPairs) Len() int {
	return len(wc)
}
func (wc WordCountPairs) Less(i, j int) bool {
	return wc[i].Count < wc[j].Count
}
func (wc WordCountPairs) Swap(i, j int) {
	wc[i], wc[j] = wc[j], wc[i]
}

// Analyze ...(Finds the number of occurances of each word ifnoring cases in a given file and sorts them.  )
func Analyze(content string) WordCountPairs {
	wordCount := make(map[string]int)

	r, _ := regexp.Compile("[^A-Za-z]+")

	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := scanner.Text()
		newString := strings.TrimSpace(strings.ToLower(r.ReplaceAllString(line, " ")))
		words := strings.Split(newString, " ")
		for _, word := range words {
			_, ok := wordCount[word]
			if ok {
				wordCount[word]++
			} else {
				wordCount[word] = 1
			}
		}
	}

	wcPairs := make(WordCountPairs, len(wordCount))

	i := int(0)
	for word, count := range wordCount {
		wcPairs[i] = Pair{word, count}
		i++
	}

	// Reverse sort
	sort.Slice(wcPairs, func(i, j int) bool {
		// Sort by word in descending order
		if wcPairs[i].Count == wcPairs[j].Count {
			return wcPairs[i].Word > wcPairs[j].Word
		}
		return wcPairs[i].Count > wcPairs[j].Count
	})

	return wcPairs
}
