package corpus

import "testing"

func TestCorpusAsFrequencyMap(t *testing.T) {
  strs := []string{`a`,`b`,`b`,`c`,`c`,`c`}
  actual := AsFrequencyMap(strs)

  expected := map[string]int{
    `a`: 1,
    `b`: 2,
    `c`: 3,
    `d`: 0,
  }
  for key, val := range expected {
    if actual[key] != val {
      t.Errorf("(%s) Expected [%d], got [%d]", key, val, actual[key])
    }
  }
}

func TestCorpusAsWordFreqs(t *testing.T) {
  counts := map[string]int{
    `a`: 1,
    `b`: 2,
    `c`: 3,
  }
  actual := AsWordFreqs(counts)
  for idx, _ := range actual {
    if actual[idx].Count != counts[actual[idx].Str] {
      str := actual[idx].Str
      ec := counts[actual[idx].Str]
      ac := actual[idx].Count
      t.Errorf("(%s) Expected Count of [%d], got [%d]", str, ec, ac)
    }
  }
}
