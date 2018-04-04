package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"sync"

	"./utils"
)

type usersMap map[string]uint32

var (
	replacer      = strings.NewReplacer(",", "", ".", "", ";", "", "?", "", "\"", "")
	users         = make(usersMap)
	usersFileName = "./users.json"
	lockUsers     = sync.RWMutex{}
	lockUsersFile = sync.Mutex{}
)

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

func serializeUsers() ([]byte, error) {
	lockUsers.RLock()
	defer lockUsers.RUnlock()

	return json.Marshal(users)
}

func persistUsers(data []byte) error {
	lockUsersFile.Lock()
	defer lockUsersFile.Unlock()

	return ioutil.WriteFile(usersFileName, data, 0644)
}

func storeUsers() {
	data, err := serializeUsers()
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to convert users: ", err)
		return
	}

	err = persistUsers(data)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to persist: ", err)
	}
}

var homeT = template.Must(template.ParseFiles("./home.html"))

func home(w http.ResponseWriter, r *http.Request) {
	err := homeT.Execute(w, users)
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed to executes: ", err)
	}
}

func updateUsers(username string) {
	username = replacer.Replace(username)
	if username != "" {
		lockUsers.Lock()
		defer lockUsers.Unlock()
		users[strings.ToLower(username)]++
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(utils.WhereAmI(), "Failed on r.ParseForm(): ", err)
		return
	}

	updateUsers(r.Form.Get("username"))
	storeUsers()
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
