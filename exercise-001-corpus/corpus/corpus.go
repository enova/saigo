package corpus

import (
	"fmt"
	"io/ioutil"
	// "sort"
	"strings"
  "regexp"
)

func PrintWords(filename string) int {
  str := loadFile(filename)
  words := prepText(str)


  f := countOccurrences(words)

  for k, v := range f {
    fmt.Printf("%[1]d %[2]s\n", v, k)
  }

  return len(f)
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

func loadFile(filename string) string {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println(err)
  }
  return string(bs)
}
