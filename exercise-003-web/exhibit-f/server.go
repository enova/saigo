package main

import (
	"html/template"
	"net/http"
)

var names map[string]int
var homeT = template.Must(template.ParseFiles("home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, names)
}

func addName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	if _, ok := names[name]; ok {
		names[name]++
	} else {
		names[name] = 1
	}
	homeT.Execute(w, names)
}

func main() {
	names = make(map[string]int)
	http.HandleFunc("/home", home)
	http.HandleFunc("/addName", addName)
	http.ListenAndServe(":8080", nil)
}
