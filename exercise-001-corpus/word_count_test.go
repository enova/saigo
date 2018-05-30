package main

import "testing"

var content string = "twinkle twinkle little star"
var wordCounts map[string]int = countWords(content)

func TestZeroCountWord(t *testing.T) {
  if wordCounts["how"] != 0 {
    t.Error("Expected 0, got ", wordCounts["how"])
  }
}

func TestSingleCountWord(t *testing.T) {
  if wordCounts["little"] != 1 {
    t.Error("Expected 1, got ", wordCounts["little"])
  }
}

func TestMulitpleCountWords(t *testing.T) {
  if wordCounts["twinkle"] != 2 {
    t.Error("Expected 2, got ", wordCounts["twinkle"])
  }
}
