package corpus

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func WordCountFromFileSorted(file string, reverse bool) (PairedList, error) {
	contents, err := StringFromFile(file)
	if err != nil {
		return nil, err
	}
	words := WordCountMapFromString(contents)
	plist := SortedMapPairedList(words, reverse)
	return plist, nil
}

func SortedMapPairedList(unsorted map[string]int, reverse bool) PairedList {
	if len(unsorted) == 0 {
		return nil
	}
	sortable := mapToPairedList(unsorted)
	if reverse {
		sort.Sort(sort.Reverse(sortable))
	} else {
		sort.Sort(sortable)
	}
	return sortable
}

func mapToPairedList(unsorted map[string]int) PairedList {
	sortable := make(PairedList, len(unsorted))
	i := 0
	for k, v := range unsorted {
		sortable[i] = Pair{k, v}
		i++
	}
	return sortable
}

func WordCountFromFile(file string) (PairedList, error) {
	contents, err := StringFromFile(file)
	if err != nil {
		return nil, err
	}
	if len(contents) > 0 {
		words := WordCountMapFromString(contents)
		wordsPL := mapToPairedList(words)
		return wordsPL, nil
	} else {
		return nil, errors.New("Empty File")
	}
}

func StringFromFile(file string) (string, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func WordCountMapFromString(str string) map[string]int {
	words := strings.Fields(str)
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	wordMap := make(map[string]int)
	for _, word := range words {
		prettyWord := strings.ToLower(reg.ReplaceAllString(word, ""))
		wordMap[prettyWord] += 1
	}
	return wordMap
}
