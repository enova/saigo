package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"./utils"
)

type usersMap map[string]uint32

var replacer = strings.NewReplacer(",", "", ".", "", ";", "", "?", "", "\"", "")
var users = make(usersMap)
var usersFileName = "./users.json"

func getUsers() {
	data, err := ioutil.ReadFile(usersFileName)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to read", usersFileName)
		return
	}

	str := string(data)
	fmt.Println(str)

	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to Unmarshal contents of", str, err)
	}
}

func storeUsers() {
	usersInJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to Unmarshal contents of users: ", err)
		return
	}
	err = ioutil.WriteFile(usersFileName, usersInJSON, 0644)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to persist contents of users: ", err)
	}
}

var homeT = template.Must(template.ParseFiles("users/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	err := homeT.Execute(w, users)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to executes: ", err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed on r.ParseForm(): ", err)
		return
	}

	username := replacer.Replace(r.Form.Get("username"))
	if username != "" {
		users[strings.ToLower(username)]++
		storeUsers()
	}
	http.Redirect(w, r, "/home", 302)
}

func main() {
	getUsers()

	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed on http.ListenAndServe(): ", err)
	}
}
