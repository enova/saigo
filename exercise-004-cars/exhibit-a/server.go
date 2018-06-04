package main

import (
	"html/template"
	"fmt"
	"net/http"
	"time"
)


var joinT *template.Template
var playT *template.Template

type Vehicle struct {
	Name  string
	Count int
}

type Vehicles struct {
	List []*Vehicle
}

type View struct {
	Username string
	Vehicles Vehicles
}

func setup(dir string) {
	joinT = template.Must(template.ParseFiles(dir + "/templates/join.html"))
	playT = template.Must(template.ParseFiles(dir + "/templates/play.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	joinT.Execute(w, nil)
}

func join(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cookie := http.Cookie{ Name: "username", Value: r.FormValue("username"), Expires: inOneYear() }
	http.SetCookie(w, &cookie)

	view := View{ Username: r.FormValue("username") }

	playT.Execute(w, view)
}

func add(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}

	//vehicleName := r.Form.Get("vehicle")

	view := View{ Username: "test" }
	fmt.Printf("Username from cookie = %v", username)
	playT.Execute(w, &view)

}

func addVehicle() {
	/*
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}
	*/

	/*
	_, ok := UserMap[userName]
	if ok {
		UserMap[userName]++
	} else {
		counter := 1
		UserMap[userName] = counter
	}
	*/
}

func exit(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	joinT.Execute(w, nil)
}


func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func poke(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", Value: "gopher", Expires: inOneYear()}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Just set cookie named 'username' set to 'gopher'")
}

func peek(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}
	fmt.Fprintf(w, "You have a cookie named 'username' set to '%s'", username)
}

func hide(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "The cookie named 'username' should be gone!")
}

func main() {
	setup(".")
	http.HandleFunc("/poke", poke)
	http.HandleFunc("/peek", peek)
	http.HandleFunc("/hide", hide)
	http.HandleFunc("/", home)
	http.HandleFunc("/join", join)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)
	http.ListenAndServe(":8080", nil)
}
