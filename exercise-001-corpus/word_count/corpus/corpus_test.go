package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	cases := []struct {
		in   string
		want PairList
	}{
		{"Are you serious? I knew you were.", PairList([]Pair{{"you", 2}})},
		{"this this this one", PairList([]Pair{{"this", 3}})},
	}

	for _, c := range cases {
		got := Analyze(c.in)
		assert.Equal(t, c.want[0].Key, got[0].Key)
		assert.Equal(t, c.want[0].Value, got[0].Value)
	}
}
