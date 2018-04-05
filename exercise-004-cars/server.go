package main

import (
  "html/template"
  "net/http"
  "runtime"
  "path/filepath"
  "time"
  "fmt"
)

var (
  playT *template.Template
  joinT *template.Template
  root string
  vmap map[string]*Vehicle
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

func inOneYear() time.Time {
  return time.Now().AddDate(1, 0, 0)
}

func setup() {
  _, thisFile, _, _ := runtime.Caller(0)
  root = filepath.Dir(thisFile)
  vmap = make(map[string]*Vehicle, 0)

  joinT = template.Must(template.ParseFiles(root + "/templates/join.html"))
  playT = template.Must(template.ParseFiles(root + "/templates/play.html"))
}

func getVehicles() Vehicles {
  fmt.Println(vmap)
  vlist := make([]*Vehicle, len(vmap))
  i := 0
  for _, vehicle := range vmap {
    if vehicle == nil {
      continue
    }
    fmt.Println("getVehicles() Adding", vehicle)
    //vlist = append(vlist, vehicle)
    vlist[i] = vehicle
    i++
  }
  fmt.Println("getVehicles()", vlist)
  fmt.Println("getVehicles() created", len(vlist), "vehicles")
  return Vehicles{List: vlist}
}

func addVehicle(model string, speed string) {
  name := model + " " + speed
  if v, e := vmap[name]; !e {
    veh := new(Vehicle)
    veh.Name = name
    veh.Count = 1
    vmap[name] = veh
  } else {
    v.Count++
    vmap[name] = v
  }
  fmt.Println(name, vmap[name])
}

func add(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  vehicle := r.Form.Get("vehicle")
  speed := r.Form.Get("speed")

  addVehicle(vehicle, speed)
  play(w, r)
}

func join(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  if len(username) != 0 {
    cookie := http.Cookie{Name: "username", Value: username, Expires: inOneYear()}
    http.SetCookie(w, &cookie)
    play(w, r)
  } else {
    joinT.Execute(w, nil)
  }
}

func play(w http.ResponseWriter, r *http.Request) {
  username, _ := r.Cookie("username")
  vehicles := getVehicles()
  fmt.Println("There are", len(vehicles.List), "vehicles")
  for _, vehicle := range vehicles.List {
    fmt.Println(vehicle.Name, vehicle.Count)
  }
  view := View{Username: username.Value, Vehicles: vehicles}
  playT.Execute(w, &view)
}

func main() {
  setup()
  http.HandleFunc("/add", add)
  http.HandleFunc("/play", play)
  http.HandleFunc("/join", join)

  fs := http.FileServer(http.Dir(root + "/public/"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.ListenAndServe(":8086", nil)
}
