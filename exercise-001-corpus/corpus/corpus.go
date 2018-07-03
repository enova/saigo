package corpus

import (
	"strings"
	"regexp"
	"sort"
)

type pair struct {
	Word string
	Count int
}

func WordCount(s string) []pair {
	// Let's clean up the string first
	r, _ := regexp.Compile("[^[[:alnum:][:space:]']+")
	s = r.ReplaceAllString(strings.ToLower(s), "")

	frequency := make(map[string]int)

	for _, w := range strings.Fields(s) {
		frequency[w] += 1
	}

	var pairs []pair

	for word, count := range frequency {
		p := pair{Word: word, Count: count }
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
