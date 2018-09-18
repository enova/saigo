package main

import (
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const DbPath = "./db.json"

var usersLock sync.RWMutex
var users = make(map[string]int)

var homeT *template.Template

func setup(dir string) {
	homeT = template.Must(template.ParseFiles(dir + "/home.html"))
}

func presentUsers() []map[string]string {
	usersLock.RLock()
	defer usersLock.RUnlock()
	presentedUsers := make([]map[string]string, 0, len(users))
	for name, count := range users {
		user := map[string]string{"name": name, "count": strconv.Itoa(count)}
		presentedUsers = append(presentedUsers, user)
	}
	sort.SliceStable(presentedUsers, func(i, j int) bool {
		nameI := presentedUsers[i]["name"]
		nameJ := presentedUsers[j]["name"]
		if users[nameI] == users[nameJ] {
			return strings.Compare(nameI, nameJ) > 0
		} else {
			return users[nameI] > users[nameJ]
		}
	})
	return presentedUsers
}

func HandleRoot(w http.ResponseWriter, _ *http.Request) {
	homeT.Execute(w, presentUsers())
}

func handlePostUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("username")
	if len(name) > 0 {
		usersLock.Lock()
		users[name]++
		usersLock.Unlock()
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		homeT.Execute(w, presentUsers())
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostUsers(w, r)
	default:
		http.Error(w, "", http.StatusNotFound)
	}
}

func Mount() {
	setup("./templates")
	RestoreState(&users, DbPath)
}

func Unmount() {
	SaveState(&users, DbPath)
}
