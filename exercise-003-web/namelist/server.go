package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeT *template.Template
var nameList = make(map[string]int)

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		homeT.Execute(w, nameList)
	case http.MethodPost:
		r.ParseForm()
		name := r.Form.Get("username")
		nameList[name]++
		homeT.Execute(w, nameList)
	default:
		fmt.Fprintln(w, "invalid request method")
	}
}

func main() {
	setup(".")
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
