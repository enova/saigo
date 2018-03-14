package corpus

import "sort"

type ByFreq []WordFreq

func (b ByFreq) Len() int {
  return len(b)
}

func (b ByFreq) Swap(i, j int) {
  b[i], b[j] = b[j], b[i]
}

func (b ByFreq) Less(i, j int) bool {
  return b[i].Count > b[j].Count
}

func SortWordFreqs(freqs []WordFreq) {
  sort.Sort(ByFreq(freqs))
}