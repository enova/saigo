package corpus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
	result1 := WordCount("Are you serious? SERIOUS, I knew you were.")
	assert.Equal(t, len(result1), 6)
	assert.Equal(t, result1[1].Word, "you")
	assert.Equal(t, result1[1].Freq, 2)
	assert.Equal(t, result1[0].Word, "serious")
	assert.Equal(t, result1[0].Freq, 2)

	result2 := WordCount(",.!")
	assert.Equal(t, len(result2), 0)
}
