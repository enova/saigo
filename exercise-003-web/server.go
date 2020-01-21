package main

import (
	"html/template"
	"net/http"
)

type Respondent struct {
	Name string
	Count int
}

type View struct {
	Respondents []Respondent
}

var view = View{[]Respondent{}}
var homeT *template.Template

func setup() {
	homeT = template.Must(template.ParseFiles("home.html"))
}

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, &view)
}

func addNameToView(name string) {
	for i, resp := range view.Respondents {
		if resp.Name == name {
			view.Respondents[i].Count++
			return
		}
	}
	view.Respondents = append(view.Respondents, Respondent{
		Name:  name,
		Count: 1,
		} )
}

func submit(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	name := r.Form.Get("name")
	addNameToView(name)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func main() {
	setup()
	http.HandleFunc("/home", home)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

