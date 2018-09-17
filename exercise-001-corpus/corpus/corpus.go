package corpus

import (
	_ "fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type Tuple struct {
	Word  string
	Count int
}

func ReadData(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func Tokenize(data string) []string {
	data = strings.ToLower(data)
	data = regexp.MustCompile("\t+").ReplaceAllString(data, " ")
	data = regexp.MustCompile("[^a-z0-9 ]").ReplaceAllString(data, "")
	return strings.Fields(data)
}

func Analyze(data string) []Tuple {
	tokenized := Tokenize(data)
	container := make(map[string]int)
	for _, v := range tokenized {
		container[v]++
	}
	return Present(container)
}

func Present(analyzed map[string]int) []Tuple {
	tuples := make([]Tuple, 0)
	for k, v := range analyzed {
		tuples = append(tuples, Tuple{k, v})
	}
	sort.SliceStable(tuples, func(i, j int) bool {
		return tuples[i].Count > tuples[j].Count
	})
	return tuples
}

func Exec(filePath string) ([]Tuple, error) {
	data, err := ReadData(filePath)
	return Analyze(data), err
}
