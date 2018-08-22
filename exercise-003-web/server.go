package main

import (
	"encoding/csv"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type View struct {
	Visitors []Visitor
}

type Visitor struct {
	Name  string
	Count int
}

type State struct {
	filename string
	visitors map[string]int
	order    []string
}

var state State = State{filename: "visitors.csv", visitors: make(map[string]int)}
var indexTemplate = template.Must(template.ParseFiles("index.html"))

func (s *State) toRecords() (records [][]string) {
	for _, name := range s.order {
		records = append(records, []string{name, strconv.Itoa(s.visitors[name])})
	}

	return
}

func (s *State) Save() {
	fileWriter, err := os.OpenFile(s.filename, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	csvWriter := csv.NewWriter(fileWriter)
	if err = csvWriter.WriteAll(s.toRecords()); err != nil {
		panic(err)
	}
}

func (s *State) Recover() {
	fileReader, err := os.Open(s.filename)
	if err != nil {
		return
	}

	csvReader := csv.NewReader(fileReader)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, val := range records {
		s.order = append(s.order, val[0])
		count, err := strconv.Atoi(val[1])
		if err != nil {
			panic(err)
		}

		s.visitors[val[0]] = count
	}
}

func (s *State) Increment(name string) {
	if _, found := s.visitors[name]; !found {
		s.order = append(s.order, name)
	}

	s.visitors[name]++
}

func (s *State) Iter() (result []Visitor) {
	for _, val := range s.order {
		result = append(result, Visitor{val, s.visitors[val]})
	}

	return
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	if err := indexTemplate.Execute(w, &View{state.Iter()}); err != nil {
		panic(err)
	}
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form is required", http.StatusBadRequest)
		return
	}

	if name := r.Form.Get("name"); name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
	} else {
		state.Increment(name)
		state.Save()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	state.Recover()

	http.HandleFunc("/submit", submit)
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8080", nil)
}
