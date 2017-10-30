package corpus

import (
  "testing"
  "github.com/stretchr/testify/assert"
)


func TestCorpus(t *testing.T) {
  //4 unique words
  result := FrequencyFactory("   Hi there! There you are.  ")
  assert.Equal(t, 4, len(result))
  assert.Equal(t, 2, result[0].Freq)
  assert.Equal(t, "are", result[1].Word)
}
