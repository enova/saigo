package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Phone information all translated to work with json functionality
type Phone struct {
    Name string `json:"name"`
		Symbol string `json:"symbol"`
		AtomicNumber int `json:"atomicnum"`
		AtomicWeight float64 `json:"atomicweight"`
		Category     string `json:"catergory"`
		Group        int `json:"group"`
		Period       int `json:"period"`
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

		// Formatting json information
		format, err := json.MarshalIndent(allPhones, "", "    ")
		if err != nil {
			panic(err)
		}

		// Printing out formatted information to webpage
		fmt.Fprintln(w, string(format))
}

func main() {
    setup()
    http.HandleFunc("/phones", phones)
		http.HandleFunc("/", phones)
    http.ListenAndServe(":8080", nil)
}
