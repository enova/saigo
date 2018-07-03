package corpus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Analyze(s string) []Pair {
	return WordCount(s)
}

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, 6, len(result))

	// Results should be ordered by Count
	assert.Equal(t, "you", result[0].Word)
	assert.Equal(t, 2, result[0].Count)

	// When the count is same it should be ordered by alphabetical order in reverse
	assert.Equal(t, "were", result[1].Word)
	assert.Equal(t, 1, result[1].Count)

	// Should clean up everything except alphanumeric and '
	assert.Equal(t, "serious", result[2].Word)

	// Should lowercase letters
	assert.Equal(t, "are", result[5].Word)
}
