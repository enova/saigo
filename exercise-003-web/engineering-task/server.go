package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var mutex = &sync.Mutex{}
var users = make(map[string]int)
var usersPath, _ = filepath.Abs("./users.txt")
var indexPath, _ = filepath.Abs("./index.html")
var indexT = template.Must(template.ParseFiles(indexPath))

func index(w http.ResponseWriter, r *http.Request) {
	indexT.Execute(w, &users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form.Get("username")
	mutex.Lock()
	users[user]++

	file, err := os.OpenFile(usersPath, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, user)
	writer.Flush()
	file.Close()
	mutex.Unlock()
	index(w, nil)
}

func readUsers() {
	file, err := os.Open(usersPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		users[scanner.Text()]++
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func main() {
	readUsers()
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", addUser)
	http.ListenAndServe(":8080", nil)
}
