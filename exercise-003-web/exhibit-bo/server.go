package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("exhibit-bo/index.gohtml"))
var users = make([]User, 1)

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
