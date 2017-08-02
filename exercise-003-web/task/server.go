package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var homeT = template.Must(template.ParseFiles("home.html"))

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/addName", addName)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//read stored names
	file, err := ioutil.ReadFile("data.txt")
	checkForErrors(err)

	people := strings.Fields(string(file))
	homeT.Execute(w, people)
}

func addName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//write new name to file
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	checkForErrors(err)
	_, err = file.WriteString(r.Form.Get("name") + " ")
	checkForErrors(err)

	http.Redirect(w, r, "/home", 303)
}

func checkForErrors(err error) {
	if err != nil {
		panic(err)
	}
}
