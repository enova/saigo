package corpus

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

// Pair is a struct to store word and its frequency
type Pair struct {
	Word string
	Freq int
}

// PairList is a slice of Pairs
type PairList []Pair

// Len is used to implement sort.Interface
func (p PairList) Len() int {
	return len(p)
}

// Less is used to implement sort.Interface
func (p PairList) Less(i, j int) bool {
	if p[i].Freq == p[j].Freq {
		return p[i].Word < p[j].Word
	}
	return p[i].Freq > p[j].Freq
}

// Swap is used to implement sort.Interface
func (p *PairList) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func splitString(s string) []string {
	reg, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, " ")
	return strings.Fields(s)
}

func sortByFreq(dict map[string]int) PairList {
	sorted := make(PairList, len(dict))
	i := 0
	for word, freq := range dict {
		sorted[i] = Pair{word, freq}
		i++
	}
	sort.Sort(&sorted)
	return sorted
}

// WordCount takes a string and returns a PariList of sorted word frequency
func WordCount(s string) PairList {
	words := splitString(s)
	dict := make(map[string]int)
	for _, word := range words {
		dict[strings.ToLower(word)]++
	}
	return sortByFreq(dict)
}
