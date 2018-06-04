package main

import (
	"html/template"
	"net/http"
)

var UserNames = make([]User, 0)
var UserMap = make(map[string]int)

var homeT *template.Template

type User struct {
	UserName string
	Counter int
}

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/exhibit-f/home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func userNames(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("userName" )
	counter := getCounter(userName)
	UserNames = append(UserNames, User{ userName, counter })
	homeT.Execute(w, UserNames)
}

func getCounter(userName string) int {
	_, ok := UserMap[userName]
	if ok {
		UserMap[userName]++
	} else {
		counter := 1
		UserMap[userName] = counter
	}

	return UserMap[userName]
}

func main() {
	setup(".")
	http.HandleFunc("/home", home)
	http.HandleFunc("/userNames", userNames)
	http.ListenAndServe(":8080", nil)
}
