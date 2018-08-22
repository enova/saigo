package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, len(result), 6)
	assert.Equal(t, result[0].Word, "you")
	assert.Equal(t, result[0].Count, 2)
}

func TestEmptyText(t *testing.T) {
	result := Analyze("")
	assert.Equal(t, len(result), 0)
}

func TestKeepAppearanceOrder(t *testing.T) {
	result := Analyze("stop words are great stop words are great")
	assert.Equal(t, len(result), 4)
	assert.Equal(t, result[0].Word, "stop")
	assert.Equal(t, result[0].Count, 2)
	assert.Equal(t, result[1].Word, "words")
	assert.Equal(t, result[1].Count, 2)
	assert.Equal(t, result[2].Word, "are")
	assert.Equal(t, result[2].Count, 2)
	assert.Equal(t, result[3].Word, "great")
	assert.Equal(t, result[3].Count, 2)
}

func TestCleansText(t *testing.T) {
	result := Analyze("???????????!golang@#!@$!@#! words???+++ and more and more")
	assert.Equal(t, len(result), 4)
	assert.Equal(t, result[0].Word, "and")
	assert.Equal(t, result[0].Count, 2)
	assert.Equal(t, result[1].Word, "more")
	assert.Equal(t, result[1].Count, 2)
	assert.Equal(t, result[2].Word, "golang")
	assert.Equal(t, result[2].Count, 1)
}
