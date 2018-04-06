package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type View struct {
	NameBook map[string]int
}

var (
	lock         sync.Mutex
	namebook     map[string]int
	homeT        = template.Must(template.ParseFiles("exercise/home.html"))
	namebookPath = "exercise/namebook.json"
)

func home(w http.ResponseWriter, r *http.Request) {
	v := View{NameBook: namebook}
	homeT.Execute(w, &v)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := strings.TrimSpace(r.Form.Get("username"))
	if username != "" {
		if err := writeNamebook(username); err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/home", 301)
}

func writeNamebook(username string) error {
	lock.Lock()
	defer lock.Unlock()
	namebook[username]++
	js, _ := json.Marshal(namebook)
	err := ioutil.WriteFile(namebookPath, js, 0644)
	return err
}

func readNamebook() error {
	lock.Lock()
	defer lock.Unlock()
	namebook = make(map[string]int)
	_, err := os.Stat(namebookPath)
	if os.IsNotExist(err) {
		var file, errCreate = os.Create(namebookPath)
		defer file.Close()
		err = errCreate
	} else if err == nil {
		namebookFile, errOpen := os.Open(namebookPath)
		if errOpen != nil {
			return errOpen
		}
		defer namebookFile.Close()
		s, errRead := ioutil.ReadFile(namebookPath)
		if errRead != nil {
			return errRead
		}
		err = json.Unmarshal(s, &namebook)
	}
	return err
}

func main() {
	err := readNamebook()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
