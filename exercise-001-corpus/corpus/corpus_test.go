package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpusShortString(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, 6, len(result))
	assert.Equal(t, "you", result[0].Word)
	assert.Equal(t, 2, result[0].Count)
}

func TestCorpusLongerString(t *testing.T) {
	result := Analyze("Are you serious? I knew you were. What is going on? I don't know. Who knows? I know. I do know.")
	assert.Equal(t, "I", result[0].Word)
	assert.Equal(t, 4, result[0].Count)
	assert.Equal(t, "know", result[1].Word)
	assert.Equal(t, 3, result[1].Count)
}
