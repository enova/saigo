package main

import (
	"html/template"
	"net/http"
	"time"
)


var joinT *template.Template
var playT *template.Template
var vehicles = Vehicles{}
var views = make([]View, 0)

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
	poke(w, r)
	username := r.FormValue("username")
	view := getView(username)
	vehicles = view.Vehicles
	playT.Execute(w, &view)
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := peek(w, r)

	view := getView(username)

	vehicle := r.Form.Get("vehicle")
	addVehicle(vehicle)
	view = updateView(username)
	playT.Execute(w, &view)
}

func addVehicle(vehicle string) {
	if !isVehicleExist(vehicle) {
		vehicles.List = append(vehicles.List, &Vehicle{ Name: vehicle, Count : 1 })
	}
}

func isVehicleExist(vehicle string) bool {
	found := false
	for _, v := range vehicles.List {
		if vehicle == v.Name {
			v.Count++
			found = true
			break
		}
	}
	return found
}

func updateView(username string) View {
	view := View {}
	for i := 0; i < len(views); i++ {
		if views[i].Username == username {
			views[i].Vehicles = vehicles
			return views[i]
		}
	}
	return view
}

func getView(username string) View {
	for i := 0; i < len(views); i++ {
		if views[i].Username == username {
			return views[i]
		}
	}
	view := View{ Username: username }
	views = append(views, view)

	return view
}

func exit(w http.ResponseWriter, r *http.Request) {
	hide(w, r)
	joinT.Execute(w, nil)
}

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func poke(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", Value: r.FormValue("username"), Expires: inOneYear()}
	http.SetCookie(w, &cookie)
}

func peek(w http.ResponseWriter, r *http.Request) string {
	username, err := r.Cookie("username")
	if err != nil {
		return ""
	}

	return username.Value
}

func hide(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
}

func main() {
	setup(".")
	http.HandleFunc("/", home)
	http.HandleFunc("/join", join)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)
	http.ListenAndServe(":8080", nil)
}
