package corpus

import (
	"strings"
	"regexp"
	"sort"
)

// Pair contains frequency of word's usage
type Pair struct {
	Word string
	Count int
}

// WordCount will return a slice of Pairs, reverse ordered by count
func WordCount(s string) []Pair {
	// Let's clean up the string first
	r, _ := regexp.Compile("[^[[:alnum:][:space:]']+")
	s = r.ReplaceAllString(strings.ToLower(s), "")

	frequency := make(map[string]int)

	for _, w := range strings.Fields(s) {
		frequency[w]++
	}

	var pairs []Pair

	for word, count := range frequency {
		p := Pair{Word: word, Count: count }
		pairs = append(pairs, p)
	}

	// In-place reverse sort
	sort.Slice(pairs, func(i, j int) bool {
		// If count is same, sort by word in descending order
		if pairs[i].Count == pairs[j].Count {
			return pairs[i].Word > pairs[j].Word
		}

		return pairs[i].Count > pairs[j].Count
	})

	return pairs
}
