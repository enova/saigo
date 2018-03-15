package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mokpro/saigo/exercise-003-web/task/visitor"
)

func home(w http.ResponseWriter, r *http.Request) {
	homeT := template.Must(template.ParseFiles("task/home.html"))
	visitorList := visitor.List()
	homeT.Execute(w, &visitorList)
}

func visit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	visitCt := visitor.Record(name)
	msg := "Hey " + name + ", the answer is 42!"
	if visitCt > 1 {
		msg += "\nSame as last " + strconv.Itoa(visitCt) + " times"
	}
	msg += "\n\n(Reference: Hitch Hiker's Guide to the Galaxy)"
	fmt.Fprintln(w, msg)
}

func main() {
	http.HandleFunc("/visit", visit)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
