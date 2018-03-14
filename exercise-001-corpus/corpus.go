package main

import (
  "os"
  "io/ioutil"
  "errors"
  "./corpus"
)

func ingestFile(filename string) (string, error) {
  buff, err := ioutil.ReadFile(filename)
  if err != nil {
    return "", errors.New("Cannot read file.")
  } else {
    return string(buff), nil
  }
}

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
    panic(err)
  }

  rawFileText, err := ingestFile(fName)
  if err != nil {
    panic(err)
  }

  wordList := corpus.GrepWords(rawFileText)
  if len(wordList) == 0 {
    return
  }

  wordFreqs := corpus.GenerateWordFreqs(wordList)
  corpus.PrintSortedWordFreqs(wordFreqs)
}