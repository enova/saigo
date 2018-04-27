package main

import (
	"html/template"
	"net/http"
	"sort"
	"time"
)

var joinT = template.Must(template.ParseFiles("templates/join.html"))
var playT = template.Must(template.ParseFiles("templates/play.html"))

var vehicles = Vehicles{}

type Vehicle struct {
	Name  string
	Count int
}

type Vehicles struct {
	List []*Vehicle
}

func (v Vehicles) Len() int           { return len(v.List) }
func (v *Vehicles) Swap(i, j int)     { v.List[i], v.List[j] = v.List[j], v.List[i] }
func (v Vehicles) Less(i, j int) bool { return v.List[i].Count < v.List[j].Count }

func (v *Vehicles) Increment(s string) {
	found := false
	for _, i := range v.List {
		if i.Name == s {
			i.Count++
			found = true
			break
		}
	}
	if !found {
		v.List = append(v.List, &Vehicle{Name: s, Count: 1})
	}
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
	case "GET":
		joinT.Execute(w, nil)
	case "POST":
		r.ParseForm()
		cookie := http.Cookie{Name: "username", Value: r.FormValue("username"), Expires: inOneYear()}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/play", 302)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", 302)
	}
	sort.Sort(&vehicles)
	v := View{Username: username.Value, Vehicles: vehicles}
	playT.Execute(w, &v)
}

func add(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", 302)
	}
	r.ParseForm()
	vehicles.Increment(r.PostFormValue("vehicle"))
	play(w, r)
}

func exit(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	joinT.Execute(w, nil)
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/exit", exit)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)
}
