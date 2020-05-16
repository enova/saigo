package main

import (
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-b/home.html")) // create a new Template and parse the template definitions from the named file

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil) // apply the parsed template to the data object while writing the output to w
}

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
