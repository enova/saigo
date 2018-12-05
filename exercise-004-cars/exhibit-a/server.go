package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Vehicle ... structure with name and count
type Vehicle struct {
	Name  string
	Count int
}

// Vehicles ... containing list of all vehicles for a particular user
type Vehicles struct {
	List []*Vehicle
}

// View ... containing username and list of vehicles associated with the user
type View struct {
	Username string
	Vehicles Vehicles
}

// Views ... Map containing username as Key and View as Value
type Views map[string]View

var allViews = make(Views)

var joinT *template.Template
var playT *template.Template

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
		if username == "" {
			fmt.Println("Username is blank")
			joinT.Execute(w, nil)
			return
		}

		cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "http://localhost:8080/play", 302)

	default:
		fmt.Println("Invalid http input method")
	}

}

func play(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}
	username := cookie.Value
	view, ok := allViews[username]

	if !ok {
		view = View{Username: username}
		allViews[username] = view
	}

	playT.Execute(w, view)
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vehicleName := r.Form.Get("vehicle")
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}
	username := cookie.Value
	addVehicleToList(username, vehicleName)

	http.Redirect(w, r, "/play", 302)

}

func addVehicleToList(username string, vehicleName string) {
	existingView := allViews[username]
	list := existingView.Vehicles.List
	found := 0

	for _, vehicle := range list {
		if vehicle.Name == vehicleName {
			vehicle.Count++
			found = 1
		}
	}
	//First time adding a vehicle to the list of vehicles
	if found == 0 {
		newVehicle := &Vehicle{vehicleName, 1}
		list = append(list, newVehicle)
	}

	existingView.Vehicles.List = list
	allViews[username] = existingView
}

func exit(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/join", 302)
}

func main() {
	setup(".")
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}
