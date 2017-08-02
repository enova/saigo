package main

import "html/template"
import "net/http"

type person struct {
	Name string
}

var storedPeople []person

var homeT = template.Must(template.ParseFiles("home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, storedPeople)
}

func addName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inputName := r.Form.Get("name")
	var inputPerson person
	inputPerson.Name = inputName
	storedPeople = append(storedPeople, inputPerson)
	http.Redirect(w, r, "/home", 303)
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/addName", addName)
	http.ListenAndServe(":8080", nil)
}
