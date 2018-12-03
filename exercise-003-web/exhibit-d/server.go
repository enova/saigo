package main

import (
	"html/template"
	"net/http"
  "fmt"
)
var loginInfo = make(map[string]int)

var homeT = template.Must(template.ParseFiles("exhibit-d/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
  username := r.Form.Get("username")

  loginInfo[username]++
	// username := r.Form.Get("username")
  homeT.Execute(w, loginInfo)

}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
