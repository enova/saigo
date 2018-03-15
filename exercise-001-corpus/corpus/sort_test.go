package corpus

import "testing"

func TestCorpusSortWordFreqs(t *testing.T) {
	freqs := []WordFreq{
		WordFreq{`they`, 5},
		WordFreq{`our`, 4},
		WordFreq{`his`, 4},
		WordFreq{`be`, 2},
		WordFreq{`am`, 2},
		WordFreq{`he`, 2},
		WordFreq{`a`, 1},
		WordFreq{`i`, 1},
		WordFreq{`us`, 1},
	}
	expected := []string{`they`, `our`, `his`, `be`, `am`, `he`, `a`, `i`, `us`}
	for idx, word := range expected {
		if freqs[idx].Str != word {
			t.Errorf("Expected [%s] to eq [%s]", freqs[idx].Str, word)
		}
	}
}
