package corpus

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int      { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

var words map[string]int

//CountWords is the kick off function
func CountWords(fileName string) map[string]int {
	readFile(fileName)
	outputWords()
	return words
}

func readFile(fileName string) bool {
	file, _ := os.Open(fileName)
	defer file.Close()

	stat, _ := file.Stat()

	bs := make([]byte, stat.Size())
	_, err := file.Read(bs)
	if err != nil {
		return false
	}
	str := string(bs)
	parseWords(str)

	return true
}

func parseWords(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	processedString := reg.ReplaceAllString(str, "")
	reg, _ = regexp.Compile("[  ]+")
	processedString = reg.ReplaceAllString(processedString, " ")
	wordList := strings.Split(processedString, " ")

	words = make(map[string]int)
	for _, value := range wordList {
		if _, ok := words[value]; ok {
			words[value]++
		} else {
			words[value] = 1
		}
	}
	return processedString
}

func outputWords() {
	pl := make(PairList, len(words))
	i := 0
	for k, v := range words {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	for _, v := range pl {
		fmt.Println(v.Key+" ", v.Value)
	}
}
