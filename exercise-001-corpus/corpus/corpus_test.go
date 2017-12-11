package corpus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
	assert := assert.New(t)

	assert.True(readFile("../7oldsamr.txt"))
	assert.Equal(parseWords("this. is. a, -test"), "this is a test")
	assert.Equal(CountWords("../7oldsamr.txt")["and"], 15)
}
