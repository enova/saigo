package main

import "testing"

func TestTransformWordCount(t *testing.T) {
  var wordCountsMap = map[string]int{
    "how": 3711,
    "i":   2138,
    "wonder": 1908,
    "what": 912,
    "you": 142,
    "are": 352,
  }
  var wordCounts = transformWordCount(wordCountsMap)
  if len(wordCounts) != 6 {
    t.Error("Expected 6, got ", len(wordCounts))
  }
}
