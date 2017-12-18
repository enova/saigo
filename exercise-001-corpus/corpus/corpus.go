package corpus

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type keyValue struct {
	key   string
	value int
}

type List []keyValue

func WordCount(filepath string) {
	file_str_lower := readFile(filepath)
	parsed_str := parseStr(file_str_lower)
	counts := countWords(parsed_str)
	sortCount(counts)
}

func readFile(filepath string) string {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	file_str := strings.ToLower(string(file))
	return strings.Replace(file_str, "\n", " ", -1)
}

func parseStr(str string) []string {
	reg, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		panic(err)
	}
	return strings.Fields(reg.ReplaceAllString(str, ""))
}

func countWords(strArr []string) map[string]int {
	word_counts := make(map[string]int)
	for _, word := range strArr {
		_, exists := word_counts[word]
		if exists {
			word_counts[word] += 1
		} else {
			word_counts[word] = 1
		}
	}
	return word_counts

}

func (l List) Len() int      { return len(l) }
func (l List) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l List) Less(i, j int) bool {
	if l[i].value != l[j].value {
		return l[i].value < l[j].value
	} else {
		return l[j].key < l[i].key
	}
}

func sortCount(wordmap map[string]int) {
	l := make(List, len(wordmap))

	i := 0
	for k, v := range wordmap {
		l[i] = keyValue{k, v}
		i++
	}

	sort.Sort(sort.Reverse(l))

	for _, keyValue := range l {
		fmt.Printf("%d %s\n", keyValue.value, keyValue.key)
	}
}
