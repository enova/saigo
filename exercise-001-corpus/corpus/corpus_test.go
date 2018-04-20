package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordCount(t *testing.T) {
	testStr := "The test is the best of the best!"
	assert.True(t, WordCount(&testStr)["the"] == 3, "should be 3")
}

func TestSortedWordCount(t *testing.T) {
	testStr := "The test is the best of the best!"
	assert.True(t, SortedWordCount(&testStr)[0] == "3 the", "should be '2 the'")
	assert.True(t, SortedWordCount(&testStr)[1] == "2 best", "should be '2 best'")
}
