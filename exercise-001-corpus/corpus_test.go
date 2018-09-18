package main

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	const data = "Are you serious? I knew you were."
	const expectedUniqueWordsCount = 6
	const targetWordType = "you"
	const targetWordCount = 2

	result := Analyze(data)
	assert.Equal(t, len(result), expectedUniqueWordsCount)
	assert.Equal(t, result[0].Word, targetWordType)
	assert.Equal(t, result[0].Count, targetWordCount)
}

func TestEmpty(t *testing.T) {
	const data = ""
	result := Analyze(data)
	assert.Empty(t, result)
}

func TestStrangeCharacters(t *testing.T) {
	const data = "&😀hello"
	const validWord = "hello"
	const validWordCount = 1
	result := Analyze(data)
	assert.Equal(t, len(result), validWordCount)
	assert.Equal(t, result[0].Word, validWord)
}
