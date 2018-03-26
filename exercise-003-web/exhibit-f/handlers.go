package main

import (
	"html/template"
	"net/http"
)

type view struct {
	Session map[string]int
}

var tpl = template.Must(template.ParseFiles("catalog.html"))

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	if len(username) != 0 {
		Session[username]++
	}

	sessionView := view{Session: Session}
	tpl.Execute(w, &sessionView)
}

// InitServerHandlers assigns handlers for each route
func InitServerHandlers() {
	http.HandleFunc("/", signup)
	http.HandleFunc("/signup", signup)
}
