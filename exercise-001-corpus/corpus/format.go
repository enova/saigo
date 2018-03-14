package corpus

import "fmt"

// prints word frequencies as Count, Str
func PrintWordFreqs(freqs []WordFreq) {
  for _, freq := range freqs {
    fmt.Println(freq.Count, freq.Str)
  }
}

// prints word frequencies as Count, Str
// in descending sorted order
func PrintSortedWordFreqs(freqs []WordFreq) {
  SortWordFreqs(freqs)
  for _, freq := range freqs {
    fmt.Println(freq.Count, freq.Str)
  }
}