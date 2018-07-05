package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var homeT *template.Template
var stats = make(map[string]int)
var mutex = &sync.Mutex{}
var quit = make(chan os.Signal, 1)
var saveFile = "./exhibit-f/saved_data"

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/exhibit-f/home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mutex.Lock()
		homeT.Execute(w, stats)
		mutex.Unlock()
	case "POST":
		r.ParseForm()
		username := r.Form.Get("username")
		mutex.Lock()
		stats[username]++
		homeT.Execute(w, stats)
		mutex.Unlock()
	default:
		fmt.Println("Invalid HTTP method")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func load() {
	data, err := ioutil.ReadFile(saveFile)

	// If we don't have file or cannot read we start fresh.
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &stats)
	check(err)
}

func save() {
	mutex.Lock()
	data, _ := json.Marshal(stats)
	mutex.Unlock()

	err := ioutil.WriteFile(saveFile, data, 0644)
	check(err)
}

func main() {
	setup(".")
	load()
	http.HandleFunc("/home", home)

	signal.Notify(quit, os.Interrupt)
	server := http.Server{Addr: ":8080", Handler: nil}
	ctx := context.Background()

	go func() {
		fmt.Println("Server is ready to handle requests")
		server.ListenAndServe()
	}()

	<-quit

	go func() {
		fmt.Println("Server	is shutting down...")
		server.Shutdown(ctx)
	}()

	fmt.Println("Server stopped")
	save()
}
