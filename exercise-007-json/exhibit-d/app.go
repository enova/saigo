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
	Name    string `json:"name"`
	Age     int    `json:"age"`
	ID      string `json:"id"`
	Carrier string `json:"carrier"`
	URL     string `json:"imageUrl"`
	Snippet string `json:"snippet"`
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
	pretty, err := json.MarshalIndent(allPhones, "", "    ")
	if err != nil {
		panic(err)
	}

	// the following command looks like garbage to human eyes
	//json.NewEncoder(w).Encode(allPhones)

	// this looks a lot better
	fmt.Fprintf(w, string(pretty))
}

func main() {
	setup()
	http.HandleFunc("/", phones)
	http.ListenAndServe(":8080", nil)
}
