package main

import (
	"html/template" // implements data-driven templates for generating HTML output safe against code injection
	"net/http"
)

type View struct { // defining new struct
	Name string
	Age  int
}

var homeT = template.Must(template.ParseFiles("exhibit-c/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	v := View{Name: "Gopher", Age: 8} // create a View struct
	homeT.Execute(w, &v)              // w = ResponseWriter, send data from &v accessible in HTML like: {{ .Name }}
}

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
