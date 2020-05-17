package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Phone ...
type Phone struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	ID       string `json:"Id"`
	ImageURL string `json:"imageUrl"`
	Snippet  string `json:"snippet"`
}

var allPhones []Phone
var db *sql.DB

func setup() {
	data, err := ioutil.ReadFile("./exhibit-d/phones.json")
	if err != nil {
		fmt.Println("Error reading phones.json")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &allPhones)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
	}

	checkDb()
	populateAllPhones()
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func connectDB() (db *sql.DB) {
	db, err := sql.Open("postgres", "dbname=phones sslmode=disable")
	panicOn(err)
	return
}

func checkDb() {
	rows, err := db.Query("SELECT COUNT(*) as count FROM phones")
	panicOn(err)
	defer rows.Close()

	if checkCount(rows) == 0 {
		seedDb()
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		panicOn(err)
	}
	return count
}

func seedDb() {
	var err error
	for _, phone := range allPhones {
		_, err = db.Exec("INSERT INTO phones(name, age, id, image_url, snippet) VALUES ($1, $2, $3, $4, $5)", phone.Name, phone.Age, phone.ID, phone.ImageURL, phone.Snippet)
		panicOn(err)
	}

}

func phones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(allPhones)
}

func populateAllPhones() {
	rows, err := db.Query("SELECT name, age, id, image_url, snippet FROM phones")
	panicOn(err)
	defer rows.Close()

	allPhones = allPhones[:0] //Clear data obtained by unmarshalling

	for rows.Next() {
		var name string
		var age int
		var id string
		var imageURL string
		var snippet string
		err := rows.Scan(&name, &age, &id, &imageURL, &snippet)
		panicOn(err)

		allPhones = append(allPhones, Phone{name, age, id, imageURL, snippet})
	}
}

func main() {
	db = connectDB()
	defer db.Close()
	setup()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe(":8080", nil)
}
