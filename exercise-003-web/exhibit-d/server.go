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
	r.ParseForm() // parses the raw query from the URL and updates r.Form
	// for POST requests, it also parses the request body as a form and puts the results into both r.PostForm and r.Form

	username := r.Form.Get("username") // get first value associated with the "username" key
	msg := "Hey " + username + ", did you try to sign-up?"
	fmt.Fprintf(w, msg) // write msg to w
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
