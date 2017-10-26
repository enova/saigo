package main

import (
	"html/template"
	"net/http"
)

//User that signs up for this great app
type User struct {
	Name        string
	Submissions int
}

var homeT = template.Must(template.ParseFiles("exhibit-f/home.html"))
var usersT = template.Must(template.ParseFiles("exhibit-f/users.html"))
var users map[string]*User

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")

	if _, ok := users[username]; ok {
		users[username].Submissions = users[username].Submissions + 1
	} else {
		users[username] = &User{username, 1}
	}

	usersT.Execute(w, &users)
}

func main() {
	users = make(map[string]*User)
	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)

}
