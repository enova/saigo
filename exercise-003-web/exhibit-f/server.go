package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var loginInfo = make(map[string]int)

var homeT *template.Template

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/exhibit-f/home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		username := r.Form.Get("username")
		if username != "" {
			loginInfo[username]++
		}
		err := homeT.Execute(w, loginInfo)
		if err != nil {
			log.Println(err)
		}
	case "GET":
		err := homeT.Execute(w, loginInfo)
		if err != nil {
			log.Println(err)
		}
	default:
		fmt.Fprint(w, "Invalid HTTP method")
	}
}

func main() {
	setup(".")
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
