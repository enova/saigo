package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

// Phone ...
type Phone struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Carrier  string `json:"carrier"`
	Id 			 string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Snippet  string `json:"snippet"`
}

var allPhones []Phone

func addFromFile(db *sql.DB) {
	data, err := ioutil.ReadFile("phones.json")
	if err != nil {
		fmt.Println("Error reading phones.json")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &allPhones)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
		os.Exit(1)
	}
	for _, phone := range allPhones {
		_, err = db.Exec("INSERT INTO phones(id, name, age, carrier, imageUrl, snippet) VALUES ($1, $2, $3, $4, $5, $6)", phone.Id, phone.Name, phone.Age, phone.Carrier, phone.ImageUrl, phone.Snippet)
		if err != nil {
			fmt.Println("Error inserting lines to db")
			os.Exit(1)
		}
	}
}

func setup() {
	db, err := sql.Open("postgres", "dbname=saigo user=postgres password=postgres sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to db")
		os.Exit(1)
	}

	rows, err := db.Query("SELECT * FROM phones")
	if err != nil {
		fmt.Println("Error running query")
		os.Exit(1)
	}
	count := 0

	for rows.Next() {
		phone := Phone{}

		err := rows.Scan(&phone.Id, &phone.Name, &phone.Age, &phone.Carrier, &phone.ImageUrl, &phone.Snippet)
		if err != nil {
			fmt.Println("Error scanning db row")
			os.Exit(1)
		}
		allPhones = append(allPhones, phone)
		count ++
	}

	if count == 0 {
		addFromFile(db)
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
