package corpus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func testWord(t *testing.T, wordsFreqList WordsFreqList, freq uint32, word string, index uint32) {
	wordFreq := wordsFreqList[index]
	assert.Equal(t, freq, wordFreq.Freq)
	assert.Equal(t, word, wordFreq.Word)
}

func TestCorpus(t *testing.T) {
	text := ""
	wordsFreqList, error := WordCounter(text)
	assert.Empty(t, wordsFreqList)
	assert.EqualError(t, error, "wordCounter: 'text' is empty")

	text = "this, that, and \n the other that. \t also, Those, THOSE? \"those\"..."
	wordsFreqList, error = WordCounter(text)
	assert.Empty(t, error)

	assert.Equal(t, 7, len(wordsFreqList))
	testWord(t, wordsFreqList, 3, "those", 0)
	testWord(t, wordsFreqList, 2, "that", 1)
	testWord(t, wordsFreqList, 1, "also", 2)
	testWord(t, wordsFreqList, 1, "and", 3)
	testWord(t, wordsFreqList, 1, "other", 4)
	testWord(t, wordsFreqList, 1, "the", 5)
	testWord(t, wordsFreqList, 1, "this", 6)
}
