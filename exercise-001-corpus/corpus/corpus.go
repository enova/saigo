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

// List type to hold an array of keyValue struct
type List []keyValue

// WordCount function called in word_count.go
func WordCount(filepath string) {
	fileStrLower := readFile(filepath)
	parsedStr := parseStr(fileStrLower)
	counts := countWords(parsedStr)
	sortCount(counts)
}

func readFile(filepath string) string {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fileStr := strings.ToLower(string(file))
	return strings.Replace(fileStr, "\n", " ", -1)
}

func parseStr(str string) []string {
	reg, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		panic(err)
	}
	return strings.Fields(reg.ReplaceAllString(str, ""))
}

func countWords(strArr []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range strArr {
		_, exists := wordCounts[word]
		if exists {
			wordCounts[word]++
		} else {
			wordCounts[word] = 1
		}
	}
	return wordCounts

}

func (l List) Len() int      { return len(l) }
func (l List) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l List) Less(i, j int) bool {
	if l[i].value != l[j].value {
		return l[i].value < l[j].value
	}
	return l[j].key < l[i].key
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
