package main

import (
	"html/template"
	"net/http"
  "io"
  "io/ioutil"
  "strings"
  "os"
  "encoding/csv"
  "log"
  "bufio"
)

type View struct {
	UserNames []string
}

var indexT = template.Must(template.ParseFiles("./index.html"))
var userNames []string

func index(w http.ResponseWriter, r *http.Request) {
	v := View{UserNames: userNames}
	indexT.Execute(w, &v)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	userNames = append(userNames, username)
  writeToFile()
	http.Redirect(w, r, "/", 301)
}

func writeToFile() {
  output := strings.Join(userNames, "\n")
  ioutil.WriteFile("names.csv", []byte(output), 0644)
}

func readInFile() {
  csvFile, _ := os.Open("names.csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  for {
      line, error := reader.Read()
      if error == io.EOF {
          break
      } else if error != nil {
          log.Fatal(error)
      }
      userNames = append(userNames,line[0])
  }
}

func main() {
  readInFile()
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
