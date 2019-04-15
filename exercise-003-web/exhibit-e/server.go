package main

import (
	"html/template"
	"net/http"
)

var homeT *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func main() {
	homeT = template.Must(template.ParseFiles("./exhibit-e/home.html"))
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
