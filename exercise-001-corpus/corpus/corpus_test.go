package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? are you? I knew you were.")

	// Number of unique words
	assert.Equal(t, 6, len(result))

	// Highest occurance of a word should be the first one in the sorted list
	assert.Equal(t, result[0].Word, "you")
	assert.Equal(t, result[0].Count, 3)

	// No alphanumeric character should be part of the word counted
	assert.Equal(t, result[2].Word, "were")
	assert.Equal(t, result[3].Word, "serious")

}
