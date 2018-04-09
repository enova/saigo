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
	Age      int    `json:"age"`
	Id       string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name     string `json:"name"`
	Snippet  string `json:"snippet"`
}

var allPhones []Phone

const (
	queryCount     = `SELECT COUNT(*) FROM phone`
	runInsertPhone = `INSERT INTO phone(age, id, imageurl, name, snippet) VALUES ($1, $2, $3, $4, $5)`
	selectPhones   = `SELECT age, id, imageurl, name, snippet FROM phone`
)

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func setup() {
	db, err := sql.Open("postgres", "dbname=saigo_json sslmode=disable")
	PanicOn(err)
	defer db.Close()
	var count int
	err = db.QueryRow(queryCount).Scan(&count)
	PanicOn(err)
	fmt.Println("number of rows in phone table: %5d", count)
	if count == 0 {
		data, err := ioutil.ReadFile("./exhibit-d/phones.json")
		if err != nil {
			fmt.Println("Error reading phones.json")
			fmt.Println(err)
			os.Exit(1)
		}
		err = json.Unmarshal(data, &allPhones)
		if err != nil {
			fmt.Println("Error in unmarshalling phones")
		}
		for _, phone := range allPhones {
			_, err := db.Exec(runInsertPhone,
				phone.Age, phone.Id, phone.ImageUrl, phone.Name, phone.Snippet)
			PanicOn(err)
		}
	} else {
		rows, err := db.Query(selectPhones)
		PanicOn(err)
		for rows.Next() {
			phone := Phone{}
			err := rows.Scan(&phone.Age, &phone.Id, &phone.ImageUrl, &phone.Name, &phone.Snippet)
			PanicOn(err)
			allPhones = append(allPhones, phone)
		}
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
