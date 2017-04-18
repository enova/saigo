package corpus

import (
    "fmt"
    "bufio"
    "sort"
    "strings"
    "os"
)

func Check(e error){
     if e != nil {
     	panic(e)
     }
}

func WordCount(file_name string){
    word_dict := make(map[string]int)

    file ,err := os.Open(file_name)
    Check(err)

    scan_this := bufio.NewScanner(file)

	var delimiters = strings.NewReplacer("\"", "", ".", "", ",", "")
	
    for scan_this.Scan() {
    	line := scan_this.Text()
		line = delimiters.Replace(line)
        words := strings.Split(line, " ")
        for i := 0; i < len(words); i++ {
			if freq, ok := word_dict[words[i]]; ok{
				word_dict[words[i]] = freq + 1
			} else {
				word_dict[words[i]] = 1
			}
        }
		delete(word_dict, "")
    }	

    n := map[int][]string{}
    var a []int
    for k, v := range word_dict {
        n[v] = append(n[v], k)
    }
    for k := range n {
        a = append(a, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
    for _, k := range a {
        for _, s := range n[k] {
            fmt.Printf("%s -> %d\n", s, k)
        }
    }
}

func main() {
     argsCLI := os.Args[1:]

     for i := 0; i < len(argsCLI); i++ {
          file_name := string(argsCLI[i])
          WordCount(file_name)
     }
}
