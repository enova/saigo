package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Phone ...
type Phone struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	ID       string `json:"id"`
	ImageURL string `json:"imageUrl"`
	Snippet  string `json:"snippet"`
}

var allPhones []Phone

func setup() {
	// Not working on the optional part to store in the DB at this point. Might get back to it later..
	data, err := ioutil.ReadFile("exhibit-d/phones.json")
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
	json.NewEncoder(w).Encode(allPhones)
}

func main() {
	setup()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe(":8080", nil)
}
