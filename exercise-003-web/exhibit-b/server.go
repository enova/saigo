package main

import (
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-b/home.html")) // create new template based on home.html

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil) // w = ResponseWriter, apply template to "nil" data
}

func main() {
	http.HandleFunc("/home", home) // routes "/home" traffic to home func
	http.ListenAndServe(":8080", nil)
}
