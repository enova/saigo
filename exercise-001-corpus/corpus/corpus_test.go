package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = "Hello there how're you today.  Hey, hello, today is today."
var wordArray = StringToArray(data)
var counter = WordCount(wordArray)
var pl = MapToSortedPL(counter)

func TestStringToArray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(len(wordArray), 10)
	assert.Equal(wordArray[2], "howre")
	assert.NotEqual(wordArray[0], "Hello")
	assert.Equal(wordArray[0], "hello")
}

func TestWordCount(t *testing.T) {
  assert := assert.New(t)

	assert.Equal(len(counter), 7)
	assert.Equal(counter["hello"], 2)
}

func TestMapToSortedPL(t *testing.T) {
	assert := assert.New(t)

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
    WordCount(wordArray)
  }
}

func BenchmarkMapToSortedPL(b *testing.B) {
  for n := 0; n < b.N; n++ {
    MapToSortedPL(counter)
  }
}
