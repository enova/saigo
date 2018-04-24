package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

var homeT *template.Template

var names = map[string]int{}

var file = "names.json"

type View struct {
	Names map[string]int
	Err   string
}

func setup() {
	homeT = template.Must(template.ParseFiles("home.html"))
	loadNames()
}

func home(w http.ResponseWriter, r *http.Request) {
	v := View{}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			v.Err = "Invalid Input"
			return
		}
		names[r.PostFormValue("name")]++
		saveNames()
	}
	v.Names = names
	homeT.Execute(w, &v)
}

func main() {
	setup()
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func loadNames() error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &names)
	return err
}

func saveNames() error {
	data, err := json.Marshal(names)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, data, 0644)
	return err
}
