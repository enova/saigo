package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime"
)

//User that signs up for this great app
type User struct {
	Name        string
	Submissions int
}

var homeT = template.Must(template.ParseFiles("exhibit-f/home.html"))
var usersT = template.Must(template.ParseFiles("exhibit-f/users.html"))
var users map[string]*User

const filename = "exhibit-f/users.gob"

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {

	//load any prior data. For simplicity, don't check for errors becuse if the file is empty it will throw and EOF error.
	load(filename, &users)

	r.ParseForm()
	username := r.Form.Get("username")

	if _, ok := users[username]; ok {
		users[username].Submissions = users[username].Submissions + 1
	} else {
		users[username] = &User{username, 1}
	}

	usersT.Execute(w, &users)

	//save the updated variable in the same file.
	err := save(filename, &users)
	check(err)
}

// save Encode via Gob to file
func save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// load Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// clears the data file
func clear(w http.ResponseWriter, r *http.Request) {
	users = make(map[string]*User)
	homeT.Execute(w, nil)

	//save the updated variable in the same file.
	err := save(filename, &users)
	check(err)
}

// check for errors
func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

func main() {
	users = make(map[string]*User)
	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/clear", clear)
	http.ListenAndServe(":8080", nil)

}
