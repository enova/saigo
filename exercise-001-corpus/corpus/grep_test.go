package corpus

import "testing"

func testWords() []string {
	str := "&a&$a $bbb b ...bb.bb.c'c^c'c c#b"
	return GrepWords(str)
}

func TestCorpusGrepWordsLen(t *testing.T) {
	words := testWords()
	if len(words) != 10 {
		t.Errorf("Expected [%d] words, got [%d]", 10, len(words))
	}
}

func TestCorpusGrepWordsMembers(t *testing.T) {
	words := testWords()
	expected := []string{`a`, `a`, `bbb`, `b`, `bb`, `bb`, `c'c`, `c'c`, `c`, `b`}
	for idx, word := range expected {
		if words[idx] != word {
			t.Errorf("Expected [%s] to eq [%s]", words[idx], word)
		}
	}
}
