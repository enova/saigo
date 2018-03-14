package corpus

import (
  "regexp"
  "strings"
)

func GrepWords(str string) []string {
  if len(str) == 0 {
    return []string{}
  }

  str = strings.ToLower(str)
  re := regexp.MustCompile(`\b[\w\']+\b`)
  return re.FindAllString(str, -1)
}