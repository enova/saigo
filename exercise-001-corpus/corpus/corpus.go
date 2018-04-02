package corpus

import (
	"sort"
	"strings"
)

type Pair struct {
	Word string
	Freq int
}

type PairList []Pair

// implement sort interface
func (this PairList) Len() int {
	return len(this)
}

// deterministic, both Freq and Word should be taken into consideration
func (this PairList) Less(i, j int) bool {
	if this[i].Freq == this[j].Freq {
		return this[i].Word < this[j].Word
	}
	return this[i].Freq < this[j].Freq
}

func (this PairList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortByFreq(dict map[string]int) PairList {
	sorted := make(PairList, len(dict))
	i := 0
	for word, freq := range dict {
		sorted[i] = Pair{word, freq}
		i++
	}
	sort.Sort(sort.Reverse(sorted))
	return sorted
}

func WordCount(s string) PairList {
	words := strings.Fields(s)
	dict := make(map[string]int)
	for _, word := range words {
		dict[strings.ToLower(word)]++
	}
	return sortByFreq(dict)
}
