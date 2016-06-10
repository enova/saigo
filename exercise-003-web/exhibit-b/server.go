package main

import (
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-b/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
