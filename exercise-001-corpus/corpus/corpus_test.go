package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testpair struct {
	value  string
	length int
	common string
	count  int
}

var tests = []testpair{
	{"Are you serious? I knew you were.", 6, "you", 2},
	{"", 1, "", 1},
	{"AAA aaa Aaa aaA", 1, "aaa", 4},
}

func TestCorpus(t *testing.T) {
	for _, pair := range tests {
		result := Analyze(pair.value)
		assert.Equal(t, len(result), pair.length)
		assert.Equal(t, result[0].Word, pair.common)
		assert.Equal(t, result[0].Count, pair.count)
	}
}
