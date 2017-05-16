package main

import (
	"html/template"
	"net/http"
	"sync"
	"os"
	"encoding/gob"
)

var homeT = template.Must(template.ParseFiles("src/webserver/home.html"))

var htmlTemplate = `{{range $key, $value := .}}
{{ $key }}: {{ $value }}
{{end}}`

func (dir NameDirectory) home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil)
}

func (dir NameDirectory) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		dir.checkUserName(username)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func (dir NameDirectory) shownames(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t := template.New("t")
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	//Read names from the file, since the file is getting updated on every insert
	decodedMap := dir.readNames()

	//Display data from the read file
	err = t.Execute(w, decodedMap)
	Check(err)
}

type NameDirectory struct {
	sync.Mutex
	nameCountMap map[string] int
}

func New() *NameDirectory {
	return &NameDirectory{
		nameCountMap: make(map[string] int),
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func (dir NameDirectory) checkUserName(username string) {

	// Lock the map so that multiple users can submit and check entries
	dir.Lock()

	// See if we can retrieve the persisted map
	decodedMap := dir.readNames()
	
	if len(decodedMap) != 0 {
		dir.nameCountMap = decodedMap
	}

	// No persisted data found, write to the regular structure
	if _, ok := dir.nameCountMap[username]; ok {

		// Name exists
		dir.nameCountMap[username]++

	} else {

		//New name
		dir.nameCountMap[username] = 1
	}

	// Persist data across sessions
	// This can be improved by not calling everytime a name is collected
	// Can be made either time based or just before the server quits or crashes
	dir.persistNames()

	// Unlock the data structures
	dir.Unlock()
}

func (dir NameDirectory) persistNames() {

	// Persist the collected names to the file
	file, err := os.OpenFile("NameCounts", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
	Check(err)
	
    e := gob.NewEncoder(file)

    // Encoding the map
    err = e.Encode(dir.nameCountMap)
	Check(err)
	file.Close()
}

func (dir NameDirectory) readNames() map[string]int{
	var decodedMap map[string]int
	file, err := os.Open("NameCounts")
	Check(err)

	// Decoding the serialized data
    d := gob.NewDecoder(file)
    err = d.Decode(&decodedMap)
	Check(err)

	//return the decoded object so that it can be used
	return decodedMap
}

func main() {
	directory := New()
	http.HandleFunc("/", directory.home)
	http.HandleFunc("/home", directory.home)
	http.HandleFunc("/signup", directory.signup)
	http.HandleFunc("/shownames", directory.shownames)
	http.ListenAndServe(":8080", nil)
}
