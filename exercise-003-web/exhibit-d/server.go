package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-d/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil) // generates home.html template, which has an input form and redirects to "/signup"
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                      // parses the form passed from home.html
	username := r.Form.Get("username") // assign "username" from parsed form data
	msg := "Hey " + username + ", did you try to sign-up?"
	fmt.Fprintf(w, msg) // uses ResponseWriter to print msg
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup) // routes "/signup" traffic to signup func
	http.ListenAndServe(":8080", nil)
}
