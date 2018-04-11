package main

import (
  "html/template"
  "net/http"
  "runtime"
  "path/filepath"
  "fmt"
)

var (
  PlayT *template.Template
  JoinT *template.Template
  root string
)

func setup() {
  _, thisFile, _, _ := runtime.Caller(0)
  root = filepath.Dir(thisFile)

  JoinT = template.Must(template.ParseFiles(root + "/templates/join.html"))
  PlayT = template.Must(template.ParseFiles(root + "/templates/play.html"))
}

func main() {
  setup()
  SetupHandlers()

  fs := http.FileServer(http.Dir(root + "/public/"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  fmt.Println("Listening on :8086...")
  http.ListenAndServe(":8086", nil)
}
