package main

import "net/http"

func add(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  AddVehicle(r.Form.Get("vehicle"), r.Form.Get("speed"))
  http.Redirect(w, r, "/play", 301)
}

func exit(w http.ResponseWriter, r *http.Request) {
  ClearUsername(w)
  http.Redirect(w, r, "/join", 301)
}

func join(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  if SetUsername(username, w) {
    http.Redirect(w, r, "/play", 301)
  } else {
    JoinT.Execute(w, nil)
  }
}

func play(w http.ResponseWriter, r *http.Request) {
  username, err := GetUsername(r)
  if(err != nil) {
    http.Redirect(w, r, "/join", 301)
  } else {
    PlayT.Execute(w, ViewVehicles(username))
  }
}

func SetupHandlers() {
  http.HandleFunc("/add", add)
  http.HandleFunc("/play", play)
  http.HandleFunc("/join", join)
  http.HandleFunc("/exit", exit)
}
