package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-f/home.html"))
var users []string

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	users = append(users, username)
	fmt.Fprint(w, "People using this wonderful form: \n")
	for i, user := range users {
		fmt.Fprintf(w, "%d. %s\n", i+1, user)
	}

}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
