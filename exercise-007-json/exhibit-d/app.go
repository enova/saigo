package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Phone ...
type Phone struct {
	Name 		string 	`json:"name"`
	Age 		int 	`json:"age"`
	ImageUrl 	string  `json:"imageUrl"`
	ID 			string  `json:"id"`
	Snippet		string  `json:"snippet"`
}

var allPhones []Phone

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func setup() {
	data, err := ioutil.ReadFile("exhibit-d/phones.json")
	PanicOn(err)

	err = json.Unmarshal(data, &allPhones)
	PanicOn(err)

}

func phones(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(allPhones)
	PanicOn(err)
}

func main() {
	setup()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe(":8080", nil)
}
