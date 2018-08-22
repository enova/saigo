package main

import (
	"fmt"
	"html/template"
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

var allUserVehicles = make(map[string]map[string]int)

func setup(dir string) {
	joinT = template.Must(template.ParseFiles(dir + "/templates/join.html"))
	playT = template.Must(template.ParseFiles(dir + "/templates/play.html"))
}

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func join(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		joinT.Execute(w, nil)
	case "POST":
		r.ParseForm()
		username := r.Form.Get("username")
		cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "http://localhost:8080/play", 301)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find a cookie named 'username'")
		return
	}

	username := cookie.Value
	userVehicles := getUserVehicles(username, allUserVehicles)
	userVehiclesView := View{username, getUserVehiclesForView(userVehicles)}
	playT.Execute(w, userVehiclesView)
}

func getUserVehicles(username string, allUserVehicles map[string]map[string]int) map[string]int {
	userVehicles, ok := allUserVehicles[username]
	if !ok {
		allUserVehicles[username] = make(map[string]int)
		userVehicles = allUserVehicles[username]
	}

	return userVehicles
}

func getUserVehiclesForView(userVehicles map[string]int) Vehicles {
	vehicles := make([]*Vehicle, len(userVehicles))
	i := 0
	for name, count := range userVehicles {
		vehicles[i] = &Vehicle{name, count}
		i++
	}

	return Vehicles{vehicles}
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vehicleName := r.Form.Get("vehicle")

	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find a cookie named 'username'")
		return
	}

	username := cookie.Value

	userVehicles := getUserVehicles(username, allUserVehicles)
	userVehicles[vehicleName]++
	http.Redirect(w, r, "http://localhost:8080/play", 301)
}

func exit(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "http://localhost:8080/join", 301)
}

func main() {
	setup(".")
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)
	http.ListenAndServe(":8080", nil)
}
