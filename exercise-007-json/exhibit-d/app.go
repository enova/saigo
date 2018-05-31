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
	Name string `json:"name"`
  Age int `json:"age"`
  Id string `json:"id"`
  Desc string `json:"snippet"`
  Image string `json:"imageUrl"`
}

var allPhones []Phone

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
	json.NewEncoder(w).Encode(allPhones)
}

func main() {
	setup()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe(":8080", nil)
}
