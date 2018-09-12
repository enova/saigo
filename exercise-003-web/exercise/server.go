package main

import (
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exercise/home.html"))
var users map[string]int

type UserList struct {
	UserTable map[string]int
}

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	users[username] = users[username] + 1
	UserList := UserList{UserTable: users}
	homeT.Execute(w, &UserList)
}

func main() {
	users = make(map[string]int)
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
