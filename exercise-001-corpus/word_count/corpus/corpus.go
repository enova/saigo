package corpus

import (
	"bufio"
	"regexp"
	"sort"
	"strings"
)

// Pair ...(Represents a word as Key and count as Value)
type Pair struct {
	Key   string
	Value int
}

// PairList ...(custom type that stores a slice of Pair (word count) representation)
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Analyze ...( a file and returns pairlist sorted according to the word count)
func Analyze(content string) PairList {
	wordCount := make(map[string]int)

	reg, _ := regexp.Compile(`[^\w\s]+`)

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := strings.TrimSpace(strings.ToLower(reg.ReplaceAllString(line, "")))
		words := strings.Split(processedLine, " ")

		for _, word := range words {
			wordCount[word]++
		}
	}

	pairList := make(PairList, len(wordCount))
	i := 0
	for word, count := range wordCount {
		pairList[i] = Pair{word, count}
		i++
	}

	sort.Sort(sort.Reverse(pairList))

	return pairList
}
