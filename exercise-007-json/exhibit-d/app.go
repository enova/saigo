package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Phone ...
type Phone struct {
	PhoneID   int       `db:"phone_id"`
	Age       int       `json:"age" db:"age"`
	Carrier   string    `json:"carrier" db:"carrier"`
	ID        string    `json:"id" db:"id"`
	ImageURL  string    `json:"imageUrl" db:"image_url"`
	Name      string    `json:"name" db:"name"`
	Snippet   string    `json:"snippet" db:"snippet"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var allPhonesJson, allPhonesDB []Phone // use 2 just to clarify not printing from json file, printing from DB data instead
var db *sqlx.DB

func setup() {
	var count int

	// set up db connection
	conn, err := sqlx.Open("postgres", "dbname=test user=postgres password=postgres sslmode=disable")
	PanicOn(err)
	db = conn

	// read json file into app
	data, err := ioutil.ReadFile("phones.json")
	if err != nil {
		fmt.Println("Error reading phones.json")
		os.Exit(1)
	}

	// break json data up into allPhones []Phone
	err = json.Unmarshal(data, &allPhonesJson)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
	}

	// create table
	initPhoneTable()

	// check if table is empty, seed if empty
	query := `SELECT COUNT(*) FROM phones`
	err = db.Get(&count, query)
	PanicOn(err)
	if count == 0 {
		seedPhoneTable(allPhonesJson)
	}
}

func main() {
	setup()
	http.HandleFunc("/phones", phones)
	http.ListenAndServe(":8080", nil)
}

func phones(w http.ResponseWriter, r *http.Request) {
	// place data from DB into allPhonesDB array
	query := `SELECT * FROM phones`
	err := db.Select(&allPhonesDB, query)
	PanicOn(err)

	// print json formatted data into browser
	output, err := json.Marshal(allPhonesDB)
	fmt.Fprintf(w, string(output))
}

func seedPhoneTable(phones []Phone) {
	for _, p := range phones {
		query := `INSERT INTO phones(age, carrier, id, image_url, name, snippet) VALUES ($1, $2, $3, $4, $5, $6)`
		_, err := db.Exec(query, p.Age, p.Carrier, p.ID, p.ImageURL, p.Name, p.Snippet)
		if err != nil {
			panic(err)
		}
	}
}

func initPhoneTable() {
	query := `
		CREATE TABLE phones(
		phone_id SERIAL PRIMARY KEY
		, id varchar(200)
		, age INTEGER
		, carrier varchar(200)
		, image_url varchar(200)
		, name varchar(200)
		, snippet varchar(200)
		, created_at TIMESTAMPTZ default now()
		, updated_at TIMESTAMPTZ default now()
		);`
	db.Exec(query)
}

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}
