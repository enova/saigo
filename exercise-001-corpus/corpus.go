package main

import (
	_ "fmt"
	"regexp"
	"sort"
	"strings"
)

type Tuple struct {
	Word  string
	Count int
}

func Analyze(data string) []Tuple {
	data = strings.ToLower(data)
	data = regexp.MustCompile("\t+").ReplaceAllString(data, " ")
	data = regexp.MustCompile("[^a-z0-9 ]").ReplaceAllString(data, "")
	tokenized := strings.Fields(data)

	container := make(map[string]int)
	for _, v := range tokenized {
		container[v]++
	}
	return Present(container)
}

func Present(analyzed map[string]int) []Tuple {
	var tuples []Tuple
	for k, v := range analyzed {
		tuples = append(tuples, Tuple{k, v})
	}
	sort.SliceStable(tuples, func(i, j int) bool {
		return tuples[i].Count > tuples[j].Count
	})
	return tuples
}
