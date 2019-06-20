package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// scanWords scans the file to parse the strings
func scanWords(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		word := strings.Replace(scanner.Text(), ".", "", -1)
		word = strings.Replace(scanner.Text(), ",", "", -1)
		word = strings.Replace(scanner.Text(), "?", "", -1)
		word = strings.Replace(scanner.Text(), ":", "", -1)
		if word[0] == '"' {
			word = word[1:]
		}
		if word[len(word)-1] == '"' {
			word = word[:len(word)-1]
		}
		words = append(words, word)
	}
	return words
}

// sortWords is the logic to reverse sort the words according to frequencies
func sortWords(words []string, m map[string]int) {
	for _, word := range words {
		m[strings.ToLower(word)] += 1
	}

	// Creating another map to swap the positions of keys and values
	n := map[int][]string{}
	var a []int

	// Filling the map
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}

	// Reverse sorting the word frequencies and their respective keys
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	// Printing the output
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
}

// Main function
func main() {

	// Checking for the necessary command line argument
	if len(os.Args) <= 1 {
		fmt.Println("Command line argument missing")
		return
	}

	// Initilizing a map
	m := make(map[string]int)

	// Scanning through the text file and adding initial info to the map
	words := scanWords(os.Args[1])

	// Sort the words by frequencies
	sortWords(words, m)
}
