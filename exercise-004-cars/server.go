package main

import (
	"net/http"
	"html/template"
	"time"
	"strings"
	"sync"
)

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

var (
	joinT        	 = template.Must(template.ParseFiles("./templates/join.html"))
	playT            = template.Must(template.ParseFiles("./templates/play.html"))
	sessionStore     map[string]map[string]int
	lock         	 sync.RWMutex
)
func inThreeMinutes() time.Time {
	return time.Now().Add(time.Minute * 3)
}

func newSession(username string) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := sessionStore[username]; !ok {
		sessionStore[username] = make(map[string]int)
	}
}

func getView(username string) View {
	lock.RLock()
	defer lock.RUnlock()
	view := View{}
	view.Username = username
	view.Vehicles.List = make([]*Vehicle, 0)
	vehicles := sessionStore[username]
	for vehicle_name, vehicle_count := range vehicles {
		vehicle := new(Vehicle)
		vehicle.Name = vehicle_name
		vehicle.Count = vehicle_count
		view.Vehicles.List = append(view.Vehicles.List, vehicle)
	}
	return view
}

func addVehicle(username string, vehicle_name string) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := sessionStore[username]; !ok {
		sessionStore[username] = make(map[string]int)
	}
	sessionStore[username][vehicle_name]++
}

func join(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		joinT.Execute(w, nil)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			joinT.Execute(w, nil)
			return
		}
		username := strings.TrimSpace(r.Form.Get("username"))
		cookie := http.Cookie{Name: "username", Value: username, Expires: inThreeMinutes()}
		http.SetCookie(w, &cookie)
		newSession(username)
		http.Redirect(w, r, "/play", 301)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/join", 301)
	} else {
		view := getView(username.Value)
		playT.Execute(w, view)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/play", 301)
	} else {
		username, err := r.Cookie("username")
		if err != nil {
			http.Redirect(w, r, "/join", 301)
		} else {
			vehicle_name := r.Form.Get("vehicle") + ":" + r.Form.Get("speed")
			addVehicle(username.Value, vehicle_name)
			http.Redirect(w, r, "/play", 301)
		}
	}
}

func main() {
	sessionStore = make(map[string]map[string]int)
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.Handle("/", http.RedirectHandler("/join", 301))
	http.Handle("/exit", http.RedirectHandler("/join", 301))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
