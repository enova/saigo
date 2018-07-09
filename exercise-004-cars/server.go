package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Vehicle stores info about given vehicle
type Vehicle struct {
	Name  string
	Count int
}

// Vehicles is a collection of vehicles
type Vehicles struct {
	List []*Vehicle
}

// HasVehicle checks whether vehicle already exists
func (v *Vehicles) HasVehicle(s string) bool {
	for _, vehicle := range v.List {
		if vehicle.Name == s {
			return true
		}
	}

	return false
}

// Add a new vehicle
func (v *Vehicles) Add(s string) {
	vehicle := Vehicle{Name: s, Count: 1}
	v.List = append(v.List, &vehicle)
}

// Increment the vehicle count
func (v *Vehicles) Increment(s string) {
	for _, vehicle := range v.List {
		if vehicle.Name == s {
			vehicle.Count++
			return
		}
	}
}

// Reset the list
func (v *Vehicles) Reset() {
	v.List = v.List[:0]
	return
}

// View holds username and list of vehicles
type View struct {
	Username string
	Vehicles Vehicles
}

var joinT = template.Must(template.ParseFiles("templates/join.html"))
var playT = template.Must(template.ParseFiles("templates/play.html"))
var vehicles = Vehicles{}

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
		http.Redirect(w, r, "/play", 302)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")

	if err != nil {
		http.Redirect(w, r, "/join", 302)
	}

	view := View{Username: username.Value, Vehicles: vehicles}

	playT.Execute(w, &view)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		name := r.Form.Get("vehicle")
		speed := r.Form.Get("speed")

		nameAndSpeed := fmt.Sprintf("%s:%s", name, speed)

		if vehicles.HasVehicle(nameAndSpeed) {
			vehicles.Increment(nameAndSpeed)
		} else {
			vehicles.Add(nameAndSpeed)
		}
	}

	http.Redirect(w, r, "/play", 302)
}

func exit(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	vehicles.Reset()
	http.Redirect(w, r, "/join", 302)
}

func main() {
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}
