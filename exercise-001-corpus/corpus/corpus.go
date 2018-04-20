package corpus

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type pair struct {
	key string
	val int
}

type sortedPairs []pair

func (p sortedPairs) Len() int           { return len(p) }
func (p sortedPairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p sortedPairs) Less(i, j int) bool { return p[i].val < p[j].val }

func (p sortedPairs) Populate(m *map[string]int) {
	i := 0
	for key, val := range *m {
		p[i].key = key
		p[i].val = val
		i++
	}
	sort.Sort(sort.Reverse(p))
}

// WordCount returns map of words and how many times they occur in the provided string
func WordCount(s *string) map[string]int {
	words := make(map[string]int)
	for _, w := range strings.Fields(*s) {
		reg, _ := regexp.Compile("[^a-zA-Z0-9']+")
		w = reg.ReplaceAllString(w, "")
		words[strings.ToLower(w)]++
	}
	return words
}

// SortedWordCount returns an array of strings with word and count which is based on the results of WordCount sorted from most common to least
func SortedWordCount(s *string) []string {
	words := WordCount(s)

	pairs := make(sortedPairs, len(words))
	pairs.Populate(&words)

	output := []string{}

	for _, p := range pairs {
		output = append(output, fmt.Sprintf("%d %s", p.val, p.key))
	}

	return output
}
