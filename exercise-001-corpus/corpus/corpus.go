package corpus

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
  "regexp"
)

func PrintWords(filename string) {
  str := loadFile(filename)
  words := prepText(str)
  f := countOccurrences(words)
  sorted := sortWords(f)
  for _, Count := range sorted {
    fmt.Printf("%d %s\n", Count.Freq, Count.Word)
  }
}

func loadFile(filename string) string {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println(err)
  }
  return string(bs)
}

func prepText(str string) []string {
   //get rid of all white space
  r := regexp.MustCompile("[\t\n\f\r ]+")
  str = r.ReplaceAllString(str, " ")

  //get rid of punctuation
  r = regexp.MustCompile("[!-/:-@[-`{-~]+")
  str = r.ReplaceAllString(str, "")

  //split string into words
  words := strings.Split(str, " ")
  for i := 0; i < len(words); i++ {
    words[i] = strings.ToLower(words[i])
  }

  return words
}

func countOccurrences(words []string) map[string]int {
  frequencies := make(map[string]int)
  for i := 1; i < len(words); i++{
    //check if word exists in map
    if v, ok := frequencies[words[i]]; ok {
      frequencies[words[i]] = v + 1
    } else {
      frequencies[words[i]] = 1
    }
  }

  return frequencies
}

type Count struct {
  Word string
  Freq int
}

func sortWords(m map[string]int) []Count {
  var ss []Count
  for k, v := range m {
      ss = append(ss, Count{k, v})
  }

  sort.Slice(ss, func(i, j int) bool {
    if ss[i].Freq == ss[j].Freq{
      return compareWords(ss[i].Word, ss[j].Word)
    }
    return ss[i].Freq > ss[j].Freq
  })

  return ss
}

func compareWords(w1, w2 string) bool {
  if w1 < w2 {
    return true
  }
  return false
}
