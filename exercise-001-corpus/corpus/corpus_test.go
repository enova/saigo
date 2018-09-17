package corpus

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, len(result), 6)
	assert.Equal(t, result[0].Word, "you")
	assert.Equal(t, result[0].Count, 2)
}

func TestEmpty(t *testing.T) {
	result := Analyze("")
	assert.Equal(t, len(result), 0)
}

func TestStrangeCharacters(t *testing.T) {
	result := Analyze("&😀hello")
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Word, "hello")
}
