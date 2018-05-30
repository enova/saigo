package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "regexp"
  "sort"
  "strings"
)

type WordCount struct {
  Key string
  Value int
}

func main() {
  file := os.Args[1]

  bs, err := ioutil.ReadFile(file)
  if err != nil {
    return
  }

  wordCounts :=  transformWordCount(countWords(scrub(string(bs))))
  sortWordCounts(wordCounts)
  print(wordCounts)
}
func scrub(content string) string {
  reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
  if err != nil {
    log.Fatal(err)
  }
  return reg.ReplaceAllString(content, "")
}

func countWords(content string) map[string]int {
  words := strings.Fields(content)
  wordCount := make(map[string]int)
  for _, word := range words {
    wordCount[word] = wordCount[word] + 1
  }
  return wordCount
}

func transformWordCount(wordCount map[string]int) []WordCount {
  var wordCounts []WordCount
  for k, v := range wordCount {
    wordCounts = append(wordCounts, WordCount{k, v})
  }
  return wordCounts
}

func sortWordCounts(wordCounts []WordCount) {
  sort.Slice(wordCounts, func(i, j int) bool {
    return wordCounts[i].Value > wordCounts[j].Value
  })
}

func print(wordCounts []WordCount) {
  for _, something := range wordCounts {
    fmt.Printf("%s => %d\n", something.Key, something.Value)
  }
}
