package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	Word  string
	Count int
}

type ByCount []Words

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func clean_string(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	reg := regexp.MustCompile("[\t\n\f\r ]+")
	str = reg.ReplaceAllString(str, " ")
	reg = regexp.MustCompile("[!-/:-@[-`{-~]+")
	str = reg.ReplaceAllString(str, "")
	return str
}

func word_list(text string) []Words {
	words := make(map[string]int)
	text_arr := strings.Split(text, " ")

	for _, word := range text_arr {
		count, ok := words[word]
		if ok {
			words[word] = count + 1
		} else {
			words[word] = 1
		}
	}

	word_arr := make([]Words, 0)
	for word, count := range words {
		word_arr = append(word_arr, Words{word, count})
	}

	return word_arr
}

func output_words(words []Words) {
	for _, word := range words {
		fmt.Println(word.Count, word.Word)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Must specify a .txt file, eg word_count test.txt")
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Cannot find file:" + args[0])
		return
	}
	defer file.Close()
	stat, _ := file.Stat()

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Cannot read file")
		return
	}

	str := string(bs)
	str = clean_string(str)
	words := word_list(str)
	sort.Sort(ByCount(words))

	output_words(words)
}
