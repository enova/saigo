package corpus

import "sort"

// sort by count
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

// sort by string value
type ByStr []WordFreq
func (b ByStr) Len() int {
  return len(b)
}
func (b ByStr) Swap(i, j int) {
  b[i], b[j] = b[j], b[i]
}
func (b ByStr) Less(i, j int) bool {
  return b[i].Str > b[j].Str
}

func SortWordFreqs(freqs []WordFreq) {
  sort.Sort(ByStr(freqs))
  sort.Sort(ByFreq(freqs))
}