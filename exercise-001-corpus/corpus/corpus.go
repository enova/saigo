package corpus

import (
	"regexp"
	"sort"
	"strings"
)

// WordFrequency keeps track of Word and number of occurances
type WordFrequency struct {
	Word  string
	Count int
}

// Analyze does something, Chris told me to write it this way!
func Analyze(initialSt string) []WordFrequency {
	filteredSt := removePuncuations(initialSt)
	wordFreqs := wordCountGenerator(filteredSt)
	return descSortOnWordCount(wordFreqs)
}

func removePuncuations(initialSt string) string {
	rg := regexp.MustCompile("[^a-zA-Z ]")
	return rg.ReplaceAllLiteralString(initialSt, "")
}

func wordCountGenerator(filteredSt string) []WordFrequency {
	wordMap := make(map[string]WordFrequency)

	for _, word := range strings.Split(filteredSt, " ") {
		_, exists := wordMap[word]

		if exists {
			wordFreq := wordMap[word]
			wordFreq.Count++
			wordMap[word] = wordFreq
		} else if word != "" {
			wordMap[word] = WordFrequency{Word: word, Count: 1}
		}
	}
	wordMapLen := len(wordMap)
	wordFreqs := make([]WordFrequency, wordMapLen)

	for _, wordFreq := range wordMap {
		wordMapLen--
		wordFreqs[wordMapLen] = wordFreq
	}

	return wordFreqs
}

func descSortOnWordCount(wordFreqs []WordFrequency) []WordFrequency {
	sort.Slice(wordFreqs, func(i, j int) bool { return wordFreqs[i].Count > wordFreqs[j].Count })
	return wordFreqs
}
