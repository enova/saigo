package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "errors"
  "./corpus"
)

// prints word frequencies in descending sorted order
func generateAndPrintWordFreqs(words []string) {
  freqs := corpus.GenerateWordFreqs(words)
  corpus.SortWordFreqs(freqs)
  for _, freq := range freqs {
    fmt.Println(freq.Count, freq.Str)
  }
}

// try to read in the provided file
func ingestFile(filename string) (string, error) {
  buff, err := ioutil.ReadFile(filename)
  if err != nil {
    return "", errors.New("Cannot read file.")
  } else {
    return string(buff), nil
  }
}

// try to get the filename from ARGV
func fileName() (string, error) {
  if len(os.Args) < 2 {
    return "", errors.New("No filename provided.")
  } else {
    return os.Args[1], nil
  }
}

func main() {
  fName, err := fileName()
  if err != nil {
    fmt.Println(err)
    return
  }

  rawFileText, err := ingestFile(fName)
  if err != nil {
    fmt.Println(err)
    return
  }

  wordList := corpus.GrepWords(rawFileText)
  if len(wordList) == 0 {
    return
  }

  generateAndPrintWordFreqs(wordList)
}