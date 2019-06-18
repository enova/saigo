package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeT = template.Must(template.ParseFiles("home.html")) // html/template lets us write the HTML in another file

var table = map[string]int{} // global table to hold usernames and counts of visits

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, nil) // writes generated HTML to ResponseWriter

	r.ParseForm()
	visitor := r.Form.Get("username") //hold current visitor name

	if count, isIn := table[visitor]; isIn {
		table[visitor] = count + 1 // user has been seen before, increment count
	} else {
		table[visitor] = 1 // new user, add to dict
	}

	for name, count := range table {
		fmt.Fprintf(w, "<b>User:</b> %v <b>Visits:</b> %v <br><br>", name, count) // print name and # of times visited
	}

}

func main() {
	http.HandleFunc("/", home)        // handle all pages with home()
	http.ListenAndServe(":8080", nil) // see web app by visiting "localhost:8080"
}
