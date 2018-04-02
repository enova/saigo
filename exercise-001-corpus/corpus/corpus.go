package corpus

import (
	"errors"
	"strings"
	"sort"
)

type WordsList []string

type WordFreqInfo struct {
	Freq uint32
	Word string
}

type WordsFreqList []WordFreqInfo

func (s WordsFreqList) Len() int {
	return len(s)
}
func (s WordsFreqList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s WordsFreqList) Less(i, j int) bool {
	return (s[j].Freq < s[i].Freq) || ((s[i].Freq == s[j].Freq) && strings.Compare(s[i].Word, s[j].Word) == -1)
}

var replacer = strings.NewReplacer(",", "", ".", "", ";", "", "?", "", "\"", "")

func WordCounter(text string) (WordsFreqList, error) {
	if (text == "") {
		return nil, errors.New("wordCounter: 'text' is empty")
	}

	wordCounts := make(map[string] uint32)
	words := strings.Fields(replacer.Replace(text))
	for _, word := range words {
		wordCounts[strings.ToLower(word)]++
	}

	wordsFreqList := make(WordsFreqList, len(wordCounts))
	i := 0
	for key, value := range wordCounts {
		wordsFreqInfo := new(WordFreqInfo)
		wordsFreqInfo.Freq = value
		wordsFreqInfo.Word = key

		wordsFreqList[i] = *wordsFreqInfo
		i++
	}

	sort.Sort(wordsFreqList)
	return wordsFreqList, nil
}