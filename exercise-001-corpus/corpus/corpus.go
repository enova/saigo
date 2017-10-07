package corpus

import (
	"fmt"
	"io/ioutil"
	// "sort"
	"strings"
  "regexp"
)

func LoadFile(filename string) int {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	str := string(bs)
  str = prepText(str)

  words := strings.Split(str, " ")

  f := countOccurrences(words)

  fmt.Println("Occurrences:")
  fmt.Println(f["off"])
  return len(words)
  // return words.count check how to do this for arrays and slices
}

func prepText(str string) string {
   //get rid of all white space
  r := regexp.MustCompile("[\t\n\f\r ]+")
  str = r.ReplaceAllString(str, " ")

  //get rid of punctuation
  r = regexp.MustCompile("[!-/:-@[-`{-~]+")
  str = r.ReplaceAllString(str, "")

  return str
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
