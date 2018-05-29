package main

import (
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("name_freq/home.html"))
var m = make(map[string]int)

func home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")

	if len(name) > 0 {
		if val, ok := m[name]; ok {
			m[name] = val + 1
		} else {
			m[name] = 1
		}
	}

	homeT.Execute(w, &m)
}

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
