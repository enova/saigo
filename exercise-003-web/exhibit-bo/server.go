package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var users = make([]User, 1)
var homeT *template.Template

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/exhibit-bo/index.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	users = append(users, User{"Bo", "Guthrie"})
	homeT.Execute(w, users)
}

func names(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	users = append(users, User{r.Form.Get("firstName"), r.Form.Get("lastName")})
	homeT.Execute(w, users)
}

func main() {
	setup("../")
	http.HandleFunc("/", index)
	http.HandleFunc("/names", names)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
