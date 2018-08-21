package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var homeT *template.Template
var userAttempts = make(map[string]int)

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/exhibit-f/home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		homeT.Execute(w, userAttempts)		
	case "POST":
		r.ParseForm()
		username := r.Form.Get("username")
		userAttempts[username]++
		homeT.Execute(w, userAttempts)
	default:
		fmt.Fprintf(w, "Not a supported method - " + r.Method)
	}
}

func main() {
	setup(".")
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}