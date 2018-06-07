package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// Phone ...
type Phone struct {
	Name     string `json:"name"`
	Snippet  string `json:"snippet`
	ImageUrl string `json:"imageUrl"`
	Carrier  string `json:"carrier"`
}

var allPhones []Phone
var phoneT = template.Must(template.ParseFiles("phone.html"))

func setup() {
	data, err := ioutil.ReadFile("phones.json")
	if err != nil {
		fmt.Println("Error reading phones.json")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &allPhones)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
	}
}

func phones(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(allPhones)
	phoneT.Execute(w, &allPhones)
}

func main() {
	setup()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe("localhost:8080", nil)
}
