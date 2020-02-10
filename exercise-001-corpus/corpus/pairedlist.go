package corpus

import (
	"fmt"
	"strings"
)

// Pair is a struct for PairedList, relating String Key to Int Value
type Pair struct {
	Key   string
	Value int
}

// PairedList is an ordered list of String, Int pairs
type PairedList []Pair

// ContainsAll checks that all elements exist in both PairedLists. Does not check for duplicates
func (p PairedList) ContainsAll(q PairedList) bool {
	for _, pair := range p {
		if !q.Contains(pair) {
			return false
		}
	}
	return true
}

// Equals checks that the pairs are the same and in the same order
func (p PairedList) Equals(q PairedList) bool {
	if len(p) != len(q) {
		return false
	}
	for i, pair := range p {
		if q[i].Value != pair.Value && q[i].Key != pair.Key {
			return false
		}
	}
	return true
}

// Count gets the count of matching pairs
func (p PairedList) Count(pair Pair) int {
	count := 0
	for _, pval := range p {
		if pair.Key == pval.Key && pair.Value == pval.Value {
			count++
		}
	}
	return count
}

// Contains checks if the supplied pair exists in the list
func (p PairedList) Contains(pair Pair) bool {
	for _, p := range p {
		if p.Key == pair.Key && p.Value == pair.Value {
			return true
		}
	}
	return false
}

func (p PairedList) String() string {
	var sb strings.Builder
	for _, pair := range p {
		sb.WriteString(fmt.Sprintf("%s %d", pair.Key, pair.Value))
	}
	return sb.String()
}

func (p PairedList) Len() int           { return len(p) }
func (p PairedList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairedList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
