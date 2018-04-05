package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}

	return original[i+1:]
}

// WhereAmI prints debug info about file, line and method name
//
func WhereAmI() string {
	function, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d %s()", chopPath(file), line, runtime.FuncForPC(function).Name())
}

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

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

type Views map[string]View

var joinT = template.Must(template.ParseFiles("templates/join.html"))
var playT = template.Must(template.ParseFiles("templates/play.html"))

var views = make(Views)

func join(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		joinT.Execute(w, nil)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(WhereAmI(), "Failed on r.ParseForm(): ", err)
			joinT.Execute(w, nil)
			return
		}

		username := r.Form.Get("username")
		if username == "" {
			fmt.Println(WhereAmI(), "username is blank")
			joinT.Execute(w, nil)
			return
		}

		cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/play", 302)
	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/play", 302)

	err := r.ParseForm()
	if err != nil {
		fmt.Println(WhereAmI(), "Failed on r.ParseForm(): ", err)
		return
	}

	vehicleName := r.Form.Get("vehicle")
	if vehicleName == "" {
		fmt.Println(WhereAmI(), "vehicleName is blank")
		return
	}

	vehicleSpeed := r.Form.Get("speed")
	if vehicleSpeed == "" {
		fmt.Println(WhereAmI(), "vehicleSpeed is blank")
		return
	}

	username, err := r.Cookie("username")
	if err != nil {
		fmt.Println(WhereAmI(), "Could not find cookie named 'username': ", err)
		return
	}

	view, exists := views[username.Value]
	if exists == false {
		fmt.Println(WhereAmI(), "Could not find view for 'username': ", username.Value)
		return
	}

	fullVehicleName := vehicleName + ":" + vehicleSpeed
	index := -1
	for i, v := range views[username.Value].Vehicles.List {
		if v.Name == fullVehicleName {
			index = i
		}
	}

	if index > -1 {
		views[username.Value].Vehicles.List[index].Count++
	} else {
		vehicle := new(Vehicle)
		vehicle.Name = fullVehicleName
		vehicle.Count = 1
		view.Vehicles.List = append(view.Vehicles.List, vehicle)
		views[username.Value] = view
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Println(WhereAmI(), "Could not find cookie named 'username': ", err)
		http.Redirect(w, r, "/join", 302)
		return
	}

	_, exists := views[username.Value]
	if exists == false {
		view := View{}
		view.Username = username.Value
		view.Vehicles.List = make([]*Vehicle, 0)
		views[username.Value] = view
	}

	playT.Execute(w, views[username.Value])
}

func exit(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/join", 302)
}

func main() {
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)

	// NOTE: this is needed to serve up .css and .js resources
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}
