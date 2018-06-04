package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, 0, len(m))
	for k, v := range m {
		p = append(p, Pair{k, v})
	}
	sort.Sort(p)
	return p
}

func output(p PairList) {
	for _, v := range p {
		fmt.Printf("%d %s \n", v.Value, v.Key)
	}
}

func normalize(str, chr string) string {
	for _, r := range chr {
		if strings.ContainsRune(str, r) {
			str = strings.Replace(str, string(r), "", -1)
		}
	}
	return str
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print("Can't read file: ", os.Args[1])
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	m := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		word = normalize(word, "\".,!?:")
		m[word] += 1
	}

	output(sortMapByValue(m))
}
