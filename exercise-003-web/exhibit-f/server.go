package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//User that signs up for this great app
type User struct {
	Name        string
	Submissions int
}

var homeT = template.Must(template.ParseFiles("exhibit-f/home.html"))
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
	// users = append(users, username)
	fmt.Fprint(w, "People using this wonderful form: \n")
	fmt.Fprint(w, "Name\tSubmissions\n")
	for _, user := range users {
		fmt.Fprintf(w, "%s\t%d\n", user.Name, user.Submissions)
	}

}

func main() {
	users = make(map[string]*User)
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)

}
