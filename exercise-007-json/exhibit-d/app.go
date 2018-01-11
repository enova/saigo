package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Phone represents a cellular phone
type Phone struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Age      int    `json:"age"`
	ImageURL string `json:"imageURL"`
	Snippet  string `json:"snippet"`
}

var allPhones []Phone

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readJSON() {
	data, err := ioutil.ReadFile("phones.json")
	if err != nil {
		fmt.Println("Error reading phones.json")
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &allPhones)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
	}
}

func phones(w http.ResponseWriter, r *http.Request) {
	allPhones := []Phone{}
	db, err := sql.Open("postgres", "dbname=postgres sslmode=disable")
	defer db.Close()
	logError(err)
	rows, err := db.Query("SELECT * FROM phones")
	logError(err)
	defer rows.Close()

	for rows.Next() {
		var (
			id       int
			name     string
			pID      string
			age      int
			imageURL string
			snippet  string
		)
		err := rows.Scan(&id, &pID, &age, &imageURL, &name, &snippet)
		logError(err)

		phone := Phone{Name: name, ID: pID, Age: age, ImageURL: imageURL, Snippet: snippet}
		allPhones = append(allPhones, phone)
	}
	json.NewEncoder(w).Encode(allPhones)
}

func setupDatabase() {
	db, err := sql.Open("postgres", "dbname=postgres sslmode=disable")
	defer db.Close()
	logError(err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM phones;").Scan(&count)
	logError(err)

	if count == 0 {
		seedDatabase()
	}
}

func seedDatabase() {
	fmt.Println("Seeding DB!")
	readJSON()

	db, err := sql.Open("postgres", "dbname=postgres sslmode=disable")
	defer db.Close()
	logError(err)

	for _, phone := range allPhones {
		_, err = db.Exec("INSERT INTO phones(age, p_id, image_url, name, snippet) VALUES ($1, $2, $3, $4, $5);", phone.Age, phone.ID, phone.ImageURL, phone.Name, phone.Snippet)
		logError(err)
	}
}

func main() {
	setupDatabase()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe(":8080", nil)
}
