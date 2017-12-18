package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// View is the data object passed into template
type View struct {
	UserNames []string
}

const fileName = "names.txt"

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
	ioutil.WriteFile(fileName, []byte(output), 0644)
}

func readInFile() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	userNames = strings.Split(string(data), "\n")
	// Remove extra line
	if userNames[len(userNames)-1] == "" {
		userNames = userNames[:len(userNames)-1]
	}
}

func main() {
	readInFile()
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
