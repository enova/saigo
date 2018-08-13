package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = "Hello there how're you today.  Hey, hello, today is today."
var dataArray = []string{"hello", "there", "howre", "you", "today", "hey", "hello", "today", "is", "today"}
var dataMap = map[string]int{"hello": 2, "there": 1, "howre": 1, "you": 1, "today": 3, "hey": 1, "is": 1}

func TestStringToArray(t *testing.T) {
	assert := assert.New(t)
	wordArray := StringToArray(data)

	assert.Equal(len(wordArray), 10)
	assert.Equal(wordArray[2], "howre")
	assert.NotEqual(wordArray[0], "Hello")
	assert.Equal(wordArray[0], "hello")
	assert.Equal(wordArray, dataArray)
}

func TestWordCount(t *testing.T) {
	assert := assert.New(t)
	counter := WordCount(dataArray)

	assert.Equal(len(counter), 7)
	assert.Equal(counter["hello"], 2)
	assert.Equal(counter, dataMap)
}

func TestMapToSortedPL(t *testing.T) {
	assert := assert.New(t)
	pl := MapToSortedPL(dataMap)

	assert.Equal(len(pl), 7)
	assert.Equal(pl[0], Pair{Key: "today", Value: 3})
	for _, pair := range pl {
		assert.NotNil(pair.Key)
		assert.NotNil(pair.Value)
	}
}

func BenchmarkStringToArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringToArray(data)
	}
}

func BenchmarkWordCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordCount(dataArray)
	}
}

func BenchmarkMapToSortedPL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MapToSortedPL(dataMap)
	}
}
