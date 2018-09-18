package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

var joinT *template.Template
var playT *template.Template

var vehicleLock sync.RWMutex

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

var usernameVehiclesByCount = make(map[string]map[string]int)

func handleJoin(w http.ResponseWriter, r *http.Request) {
	username, _ := r.Cookie("username")
	if isCookieValueSet(username) {
		http.Redirect(w, r, "/play", http.StatusTemporaryRedirect)
	} else if r.Method == http.MethodPost {
		handleCreateUser(w, r)
	} else {
		joinT.Execute(w, nil)
	}
}
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	if len(username) > 0 {
		cookie := buildCookie("username", r.Form.Get("username"))
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/play", http.StatusFound)
	} else {
		joinT.Execute(w, nil)
	}

}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	username, _ := r.Cookie("username")
	if !isCookieValueSet(username) {
		http.Redirect(w, r, "/join", http.StatusFound)
	} else if r.Method == http.MethodPost {
		handleCreateVehicle(w, r)
	} else {
		playT.Execute(w, presentGame(username.Value))
	}
}

func handleCreateVehicle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vehicleName, err := buildVehicleName(r.Form.Get("vehicle"), r.Form.Get("speed"))
	if err != nil {
		http.Redirect(w, r, "/play", http.StatusFound)
		return
	}
	username, _ := r.Cookie("username")

	vehicleLock.Lock()

	if usernameVehiclesByCount[username.Value] == nil {
		usernameVehiclesByCount[username.Value] = make(map[string]int)
		usernameVehiclesByCount[username.Value][vehicleName]++
	} else {
		usernameVehiclesByCount[username.Value][vehicleName]++
	}
	vehicleLock.Unlock()

	playT.Execute(w, presentGame(username.Value))
}

func handleExit(w http.ResponseWriter, r *http.Request) {
	cookie := buildCookie("username", "")
	cookie.MaxAge = -1
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/join", http.StatusFound)
}

func presentGame(username string) interface{} {
	vehicleLock.RLock()
	defer vehicleLock.RUnlock()
	view := new(View)
	view.Username = username

	for name, count := range usernameVehiclesByCount[username] {
		vehicle := Vehicle{Name: name, Count: count}
		view.Vehicles.List = append(view.Vehicles.List, &vehicle)
	}
	sort.SliceStable(view.Vehicles.List, func(i, j int) bool {
		nameI := view.Vehicles.List[i].Name
		nameJ := view.Vehicles.List[j].Name
		return strings.Compare(nameI, nameJ) < 0
	})
	return view
}

func buildVehicleName(vehicle string, speed string) (string, error) {
	if len(vehicle) == 0 || len(speed) == 0 {
		return "", errors.New("Invalid form params")
	} else {
		return fmt.Sprintf("%s: %s", vehicle, speed), nil
	}
}

func setup(dir string) {
	joinT = template.Must(template.ParseFiles(dir + "/join.html"))
	playT = template.Must(template.ParseFiles(dir + "/play.html"))
}

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func isCookieValueSet(cookie *http.Cookie) bool {
	return cookie != nil && len(cookie.Value) > 0
}

func buildCookie(name string, value string) http.Cookie {
	return http.Cookie{Name: name, Value: value, Expires: inOneYear()}
}

func main() {
	setup("./templates")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))

	http.HandleFunc("/", handleJoin)
	http.HandleFunc("/join", handleJoin)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/exit", handleExit)
	http.ListenAndServe(":8080", nil)
}
