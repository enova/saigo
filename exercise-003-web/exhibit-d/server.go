package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-d/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	msg := "Hey " + username + ", did you try to sign-up?"
	fmt.Fprintf(w, msg)
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
