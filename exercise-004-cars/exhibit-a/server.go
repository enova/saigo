package main

import (
	"fmt"
  "html/template"
	"net/http"
	"time"
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

var joinT = template.Must(template.ParseFiles("./templates/join.html"))
var playT = template.Must(template.ParseFiles("./templates/play.html"))

var views []View = make([]View, 0)

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

func join(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
  http.SetCookie(w, &cookie)
  playT.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
  playT.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
  joinT.Execute(w, nil)
}
func add(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username, err := r.Cookie("username")
  if err != nil {
    fmt.Fprintf(w, "Could not find cookie named 'username'")
    return
  }
  vehicleName := r.Form.Get("vehicle")
  view := getView(username.Value)
  count := getCount(vehicleName, view)
  vehicle := Vehicle{vehicleName, count + 1}
  vehicles := updateVehicles(view, vehicle)
  view = View{username.Value, Vehicles{vehicles}}
  updateViews(view)
  playT.Execute(w, &view)
}

func updateViews(view View) {
  fmt.Println("Views: %s \n", views)
  index := 0
  for i:=0;i<len(views);{
    v := views[i]
    if v.Username == view.Username {
      index = i
    }
    i = i + 1
  }
  if index == 0 {
    views = append(views, view)
  } else {
    views[index] = view
  }
}

func updateVehicles(view View, vehicle Vehicle) []*Vehicle {
  vehicles := view.Vehicles.List
  for i:=0;i<(len(vehicles));{
    v := vehicles[i]
    if v != nil && v.Name == vehicle.Name {
      vehicles[i] = &vehicle
      return vehicles
    }
    i = i + 1
  }
  vehicles = append(vehicles, &vehicle)
  return vehicles
}

func getView(username string) View {
  for i := 0; i < len(views); {
    if views[i].Username == username {
      return views[i]
    }
    i = i + 1
  }
  vehicles := make([]*Vehicle, 0)
  return View{username, Vehicles{vehicles}}
}

func getCount(vehicleName string, view View) int {
  vehicles := view.Vehicles.List
  for i := 0; i < (len(vehicles)); {
    vehicle := vehicles[i]
    if vehicle != nil && vehicle.Name == vehicleName {
      return vehicles[i].Count
    }
    i = i + 1
  }
  return 0
}

func main() {
	http.HandleFunc("/poke", poke)
	http.HandleFunc("/peek", peek)
	http.HandleFunc("/hide", hide)
  http.HandleFunc("/join", join)
  http.HandleFunc("/play", play)
  http.HandleFunc("/add", add)
  http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
