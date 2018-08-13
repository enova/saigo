package corpus

import (
	"sort"
	"strings"
	"unicode"
)

// StringToArray turns a string into an array of words from that string
func StringToArray(data string) []string {
	strData := strings.ToLower(data)
	strData = strings.Replace(strData, "'", "", -1)

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(strData, f)
}

// WordCount take in an array of strings and return a map[string]int of count of words
func WordCount(data []string) map[string]int {
	counter := make(map[string]int)
	for _, word := range data {
		counter[word]++
	}
	return counter
}

// MapToSortedPL returns the map's key/value as a Pair, sorted
func MapToSortedPL(data map[string]int) PairList {
	pl := make(PairList, len(data))
	i := 0
	for k, v := range data {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Pair {Key string, Value int}
type Pair struct {
	Key   string
	Value int
}

// PairList is an array of Pair types
type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value { // makes the result deterministic
		return p[i].Key < p[j].Key
	}
	return p[i].Value < p[j].Value
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
