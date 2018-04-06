package corpus

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func (p wordCountList) Len() int {
	return len(p)
}

func (p wordCountList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p wordCountList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type wordCountPair struct {
	Key   string
	Value int
}

type wordCountList []wordCountPair

func cleanFileContents(contents string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := strings.TrimSpace(reg.ReplaceAllString(contents, " "))
	return processedString, err
}

func countWords(wordSlice []string) (sortedList wordCountList) {
	wordMap := make(map[string]int)

	for _, word := range wordSlice {
		if val, ok := wordMap[word]; ok {
			wordMap[word] = val + 1
		} else {
			wordMap[word] = 1
		}
	}
	i := 0
	sortedList = make(wordCountList, len(wordMap))
	for k, v := range wordMap {
		sortedList[i] = wordCountPair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(sortedList))
	return
}

func formatOutput(wordCounts wordCountList) (output string) {
	for i := range wordCounts {
		output += fmt.Sprintf("\n%d %s", wordCounts[i].Value, wordCounts[i].Key)
	}
	return
}

// ParseFile reads the file contents of a passed in file, cleans them of punctuation,
//  runs them through a word counting function, and then formats the output to
//  return to word_count.go
func ParseFile(file string) (string, error) {
	var output string
	fileContents, err := ioutil.ReadFile(file)
	cleanedString, err := cleanFileContents(string(fileContents))
	if err == nil {
		wordCounts := countWords(strings.Fields(cleanedString))
		output = formatOutput(wordCounts)
	}
	return output, err
}

func main() {
}
