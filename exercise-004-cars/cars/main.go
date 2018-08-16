package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var joinT *template.Template
var playT *template.Template
var vehicles = Vehicles{}

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func setup(dir string) {
	joinT = template.Must(template.ParseFiles(dir + "/join.html"))
	playT = template.Must(template.ParseFiles(dir + "/play.html"))
}

func join(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		joinT.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		name := r.Form.Get("username")
		cookie := http.Cookie{Name: "username", Value: name, Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/play", http.StatusSeeOther)
	default:
		fmt.Fprintln(w, "invalid request method")
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		username, err := r.Cookie("username")
		if err != nil {
			fmt.Fprintf(w, "Could not find cookie named 'username'")
		}

		view := View{Username: username.Value, Vehicles: vehicles}
		playT.Execute(w, &view)
	default:
		fmt.Fprintln(w, "invalid request method")
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.ParseForm()
		name := r.Form.Get("vehicle")
		speed := r.Form.Get("speed")
		vehicle := Vehicle{Name: name, Speed: speed, Count: 1}
		vehicles.Add(vehicle)

		http.Redirect(w, r, "/play", http.StatusSeeOther)
	default:
		fmt.Fprintln(w, "invalid request method")
	}
}

func main() {
	setup("templates")
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)
}

// Vehicle info
type Vehicle struct {
	Name  string
	Speed string
	Count int
}

// Vehicles array of vehicles
type Vehicles struct {
	List []*Vehicle
}

// View things necessary to render on html
type View struct {
	Username string
	Vehicles Vehicles
}

// Add Vehicle to Vehicles List or increase the count if it already exists
func (v *Vehicles) Add(newV Vehicle) {
	for _, vehicle := range v.List {
		if vehicle.Name == newV.Name && vehicle.Speed == newV.Speed {
			vehicle.Count++
			return
		}
	}
	v.List = append(v.List, &newV)
}
