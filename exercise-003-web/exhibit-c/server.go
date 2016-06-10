package main

import (
	"html/template"
	"net/http"
)

type View struct {
	Name string
	Age  int
}

var homeT = template.Must(template.ParseFiles("exhibit-c/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	v := View{Name: "Gopher", Age: 8}
	homeT.Execute(w, &v)
}

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
