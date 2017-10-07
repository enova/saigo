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
  printSorted(f)
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

func printSorted(m map[string]int) {
  type kv struct {
          Key   string
          Value int
      }

  var ss []kv
  for k, v := range m {
      ss = append(ss, kv{k, v})
  }

  sort.Slice(ss, func(i, j int) bool {
    if ss[i].Value == ss[j].Value{
      return compareWords(ss[i].Key, ss[j].Key)
    }
    return ss[i].Value > ss[j].Value

  })

  // return ss
  for _, kv := range ss {
        fmt.Printf("%d %s\n", kv.Value, kv.Key)
  }
}

func compareWords(w1, w2 string) bool {
  if w1 < w2 {
    return true
  }
  return false
}
