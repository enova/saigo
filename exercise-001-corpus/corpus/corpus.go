package corpus

import (
	"sort"
	"strings"
	"unicode"
)

type Pair struct {
	Word  string
	Count int
}

func onlyCharactersAndSpaces(r rune) rune {
	if (r >= 'a' && r <= 'z') || unicode.IsSpace(r) {
		return r
	}

	return -1
}

func tokenize(text string) []string {
	text = strings.ToLower(text)
	text = strings.Map(onlyCharactersAndSpaces, text)
	return strings.Fields(text)
}

func Analyze(text string) []Pair {
	tokens := tokenize(text)
	wordCountMap := make(map[string]int)
	var appearanceOrder []string

	for _, word := range tokens {
		count, found := wordCountMap[word]
		wordCountMap[word] = count + 1
		if !found {
			appearanceOrder = append(appearanceOrder, word)
		}
	}

	wordCount := make([]Pair, len(wordCountMap))

	for i, word := range appearanceOrder {
		wordCount[i] = Pair{word, wordCountMap[word]}
	}

	sort.SliceStable(wordCount, func(i, j int) bool {
		return wordCount[i].Count > wordCount[j].Count
	})

	return wordCount
}
