package server

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var joinT = template.Must(template.ParseFiles("./templates/join.html"))
var playT = template.Must(template.ParseFiles("./templates/play.html"))

var vehicles = Vehicles{}

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

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func join(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		joinT.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		username := r.Form.Get("username")
		cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/play", 307)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
	}

	view := View{Username: username.Value, Vehicles: vehicles}
	playT.Execute(w, &view)
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vehicleName := r.Form.Get("vehicle")
	vehicle := Vehicle{Name: vehicleName, Count: 1}
	vehicles.updateCount(vehicle)

	http.Redirect(w, r, "/play", 307)
}

func (vehicles *Vehicles) updateCount(vehicle Vehicle) {
	for _, v := range vehicles.List {
		if v.Name == vehicle.Name {
			v.Count++
			return
		}
	}
	vehicles.List = append(vehicles.List, &vehicle)
}

func main() {
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)
}
