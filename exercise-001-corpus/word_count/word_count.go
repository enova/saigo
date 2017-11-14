package wordCount

import (
	"regexp"
	"sort"
	"strings"
)

//Words creates a word struct to store a word and how many times it appears
type Word struct {
	Key   string
	Count int
}

//ByCount sorts all words by how many times they appear
type ByCount []Word

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func cleanString(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	reg := regexp.MustCompile("[\t\n\f\r ]+")
	str = reg.ReplaceAllString(str, " ")
	reg = regexp.MustCompile("[!-/:-@[-`{-~]+")
	str = reg.ReplaceAllString(str, "")
	return str
}

func sortWords(wordCounts map[string]int) []Word {
	wordSlice := make([]Word, 0)

	for word, count := range wordCounts {
		wordSlice = append(wordSlice, Word{word, count})
	}

	sort.Sort(ByCount(wordSlice))
	return wordSlice
}

func wordList(text string) map[string]int {
	wordCounts := make(map[string]int)
	wordArr := strings.Split(text, " ")

	for _, v := range wordArr {
		wordCounts[v] = wordCounts[v] + 1
	}

	return wordCounts
}

func Count(str string) []Word {
	str = cleanString(str)
	wordCounts := wordList(str)

	return sortWords(wordCounts)
}
