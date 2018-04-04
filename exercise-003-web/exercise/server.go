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
		lock.Lock()
		namebook[username]++
		js, _ := json.Marshal(namebook)
		err := writeNamebook(js)
		lock.Unlock()
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/home", 301)
}

func writeNamebook(j []byte) error {
	err := ioutil.WriteFile(namebookPath, j, 0644)
	return err
}

func readNamebook() error {
	namebook = make(map[string]int)
	_, err := os.Stat(namebookPath)
	if os.IsNotExist(err) {
		var file, errCreate = os.Create(namebookPath)
		file.Close()
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
	lock.Lock()
	err := readNamebook()
	lock.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
