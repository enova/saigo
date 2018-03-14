package main

import (
  "sort"
  "fmt"
)

type WordFreq struct {
  Str string
  Count int
}

type ByFreq []WordFreq

func (b ByFreq) Len() int {
  return len(b)
}

func (b ByFreq) Swap(i, j int) {
  b[i], b[j] = b[j], b[i]
}

func (b ByFreq) Less(i, j int) bool {
  return b[i].Count > b[j].Count
}

// converts a []string to a map[string]int containing
// counts of each string in words
func AsFrequencyMap(words []string) map[string]int {
  counts := make(map[string]int)
  for _, word := range words {
    counts[word] += 1
  }
  return counts
}

// converts a map[string]int to a []WordFreq
func AsWordFreqs(counts map[string]int) []WordFreq {
  wordFreqs := make([]WordFreq, 0, len(counts))
  for word := range counts {
    wordFreqs = append(wordFreqs, WordFreq{word, counts[word]})
  }
  return wordFreqs
}

func PrintWordFreqs(words []string) {
  wordFreqs := AsWordFreqs(AsFrequencyMap(words))
  sort.Sort(ByFreq(wordFreqs))
  for _, freq := range wordFreqs {
    fmt.Println(freq.Count, freq.Str)
  }
}