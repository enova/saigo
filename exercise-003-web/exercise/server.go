package main

import (
	"bufio"
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// View is the data object passed into template
type View struct {
	UserNames map[string]int
}

var (
	indexT    = template.Must(template.ParseFiles("./index.html"))
	userNames map[string]int
)

func index(w http.ResponseWriter, r *http.Request) {
	v := View{UserNames: userNames}
	indexT.Execute(w, &v)
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	userNames = updateWordAcount(username, userNames)
	error := writeUsernamesToFile(userNames, "names.csv")
	if error != nil {
		log.Fatal(error)
	}
	http.Redirect(w, r, "/", 301)
}

func writeUsernamesToFile(userNames map[string]int, fileName string) error {
	file, _ := os.Create("names.csv")
	defer file.Close()
	w := csv.NewWriter(file)
	var err error
	for key, value := range userNames {
		record := []string{key, strconv.Itoa(value)}
		if e := w.Write(record); e != nil {
			err = e
		}
	}
	w.Flush()
	return err
}

func readUsernameFile(filename string) (map[string]int, error) {
	userNames = make(map[string]int)
	csvFile, _ := os.Open("names.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var err error
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		count, error := strconv.Atoi(line[1])
		if error != nil {
			err = error
		}
		userNames[line[0]] = count
	}

	return userNames, err
}

func updateWordAcount(s string, words map[string]int) map[string]int {
	var wordExists = words[s]
	if wordExists != -1 {
		words[s] = words[s] + 1
	} else {
		words[s] = 1
	}
	return words
}

func main() {
	tempUserNames, error := readUsernameFile("names.csv")
	if error != nil {
		log.Fatal(error)
	}
	userNames = tempUserNames
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
