package corpus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpus(t *testing.T) {
	assert := assert.New(t)

	strArr := []string{"are", "you", "serious", "i", "knew", "you", "were"}
	result := countWords(strArr)

	assert.Equal(readFile("../corpus/test.txt"), "are you serious? i knew you were. ")
	assert.Equal(parseStr("are you serious? i knew you were."), strArr)
	assert.Equal(result["you"], 2)
}
