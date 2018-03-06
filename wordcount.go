package main

import "fmt"
import "os"
import "reflect"
import "sort"
import "strings"

// Optimization 1: used a fixed sized byte buffer to stream parts of a file in at a time
// Optimization 2: don't convert the file content to an array of words, directly insert into the map of string to int
func main() {
  filename := os.Args[1]
  fmt.Printf("%v given as filename arg\n", filename)

  file, err := os.Open(filename)
  defer file.Close()
  if err != nil { panic(err) }

  // I discovered I can use ioutil.ReadFile
  stats, _ := os.Stat(filename)
  filesize := stats.Size()
  //fmt.Println(reflect.TypeOf(stats))
  fmt.Printf("file size: %v bytes\n", filesize)

  filecontent := make([]byte, filesize)
  count, err2 := file.Read(filecontent)
  if err2 != nil { panic(err2) }
  fmt.Printf("read %v bytes\n", count)

  //wordcount := map[string]int{"r": 3, "b": 2, "a": 3, "aslsdf": 33, "dude":4, "werd": 11 }
  wordcount := countwords(string(filecontent))
  printwordcount(wordcount)
}

func countwords(fc string) map[string]int {
  arrwords := strings.Fields(fc)
  wcounter := make(map[string]int)
  for _, str := range arrwords {
    wcounter[str]++
  }
  fmt.Println(wcounter)
  return wcounter
}

func printwordcount(wc map[string]int) {
  wcs := make([]wordcnt, len(wc))
  var i int
  for k, v := range wc {
    //wcs = append(wcs, wordcnt{count: v, word: k})
    wcs[i] = wordcnt{count: v, word: k}
    i++
  }
  //for _, v := range wcs { fmt.Printf("%v  %v\n", v.count, v.word) }

  wcss := wordcntslc(wcs)
  sort.Sort(&wcss)
  for _, v := range wcss {
    fmt.Printf("%v  %v\n", v.count, v.word)
  }
}

type wordcnt struct { count int; word string }

type wordcntslc []wordcnt

func (w *wordcntslc) Len() int { return len(*w) }
func (w *wordcntslc) Swap(i, j int) { (*w)[i], (*w)[j] = (*w)[j], (*w)[i] }
func (w *wordcntslc) Less(i, j int) bool { return (*w)[i].count < (*w)[j].count }
