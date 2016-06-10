package main

import (
	"html/template"
	"net/http"
)

var homeT *template.Template

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/exhibit-e/home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func main() {
	setup(".")
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
