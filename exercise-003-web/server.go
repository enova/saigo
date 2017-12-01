package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

var users map[string]int

var homeT = template.Must(template.ParseFiles("home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("users.json")
	_ = json.Unmarshal(data, &users)
	homeT.Execute(w, &users)
}

func addName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	data, _ := ioutil.ReadFile("users.json")
	_ = json.Unmarshal(data, &users)
 	users[name] = users[name] + 1
	b, _ := json.Marshal(users)
	_ = ioutil.WriteFile("users.json", b, 0644)
	homeT.Execute(w, &users)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/addname", addName)
	http.ListenAndServe(":8080", nil)
}
