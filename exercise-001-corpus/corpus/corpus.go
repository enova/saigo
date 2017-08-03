package corpus

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Element is a struct for storing words and their frequencies in a string
type Element struct {
	Word  string
	Count int
}

// ByCount is an array of Elements for sorting
type ByCount []Element

func (bc ByCount) Len() int {
	return len(bc)
}

func (bc ByCount) Less(i, j int) bool {
	return bc[i].Count < bc[j].Count
}

func (bc ByCount) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}

// Analyze akes a string, sanitizes it, counts the frequencies of words, and sorts them in descending order by frequency
func Analyze(str string) {
	str = sanitize(str)
	strs := count(str)
	bc := sortCount(strs)
	for i := 0; i < len(bc); i++ {
		fmt.Println(bc[i].Count, " ", bc[i].Word)
	}
}

func sanitize(str string) string {
	str = strings.ToLower(str)
	exp, _ := regexp.Compile("[^A-Za-z0-9']")
	str = exp.ReplaceAllString(str, " ")
	exp, _ = regexp.Compile("[\t\n\v\f\r ]+")
	str = exp.ReplaceAllString(str, " ")
	return str
}

func count(str string) map[string]int {
	strs := strings.Split(str, " ")
	count := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		count[strs[i]]++
	}
	return count
}

func sortCount(count map[string]int) ByCount {
	bc := make(ByCount, len(count))
	i := 0
	for w, c := range count {
		bc[i] = Element{w, c}
		i++
	}
	sort.Sort(sort.Reverse(bc))
	return bc
}
