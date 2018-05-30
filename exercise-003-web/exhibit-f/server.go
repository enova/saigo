package main

import (
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("./exhibit-f/home.html"))
var userLogins map[string]int

type MapContainer struct {
  UserLogins map[string]int
}

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  userLogins[username] = userLogins[username] + 1
  mapContainer := MapContainer{UserLogins: userLogins}
  homeT.Execute(w, &mapContainer)
}

func main() {
  userLogins = make(map[string]int)
	http.HandleFunc("/home", home)
  http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
