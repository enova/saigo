package corpus

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func WordCount(file_name string) {
	counts := make(map[string]int)

	file, err := os.Open(file_name)
	Check(err)

	scanner := bufio.NewScanner(file)

	var delimiters = strings.NewReplacer("\"", "", ".", "", ",", "", "?", "")

	for scanner.Scan() {
		line := scanner.Text()
		line = delimiters.Replace(line)
		words := strings.Fields(line)
		for _, word := range words {
			if _, ok := counts[word]; ok {
				counts[word]++
			} else {
				counts[word] = 1
			}
		}
	}

	sortedWords := sortByWordCount(counts)
	for _, v := range sortedWords {
		fmt.Println(v)
	}
}

type Count struct {
	word string
	freq int
}

type lessFunc func(p1, p2 *Count) bool

type multiSorter struct {
	counts []Count
	less   []lessFunc
}

func (ms *multiSorter) Sort(counts []Count) {
	ms.counts = counts
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.counts)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.counts[i], ms.counts[j] = ms.counts[j], ms.counts[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.counts[i], &ms.counts[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

type ByFreq []Count

func sortByWordCount(wordCounts map[string]int) ByFreq {
	words := func(c1, c2 *Count) bool {
		return c1.word < c2.word
	}

	freq := func(c1, c2 *Count) bool {
		return c1.freq > c2.freq
	}

	b := make(ByFreq, len(wordCounts))
	i := 0

	for k, v := range wordCounts {
		b[i] = Count{k, v}
		i++
	}

	OrderedBy(freq, words).Sort(b)
	return b
}

func main() {
	argsCLI := os.Args[1:]

	for i := 0; i < len(argsCLI); i++ {
		file_name := string(argsCLI[i])
		WordCount(file_name)
	}
}
