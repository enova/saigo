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
	Name string `json:"name"`
	Age int `json:"age"`
	Id string `json:"id"`
	Carrier string `json:"carrier"`
	ImageURL string `json:"imageUrl"`
	Snippet string `json:"snippet"`
}

var allPhones []Phone

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func setup() {
	db, db_err := sql.Open("postgres", "dbname=phones sslmode=disable")
	PanicOn(db_err)

	var have_data bool
	sel_err := db.QueryRow("select exists( select * from phones)").Scan(&have_data)
	PanicOn(sel_err)
	if !have_data {
		data, err := ioutil.ReadFile("phones.json")
		if err != nil {
			fmt.Println("Error reading phones.json",err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(data, &allPhones)
		if err != nil {
			fmt.Println("Error in unmarshalling phones")
		}

		for _, phone := range allPhones {
			_, ins_err := db.Exec("Insert into phones(name,age,id,carrier,imageurl,snippet) VALUES ($1, $2, $3, $4, $5, $6)",
				phone.Name,
				phone.Age,
				phone.Id,
				phone.Carrier,
				phone.ImageURL,
				phone.Snippet)
			PanicOn(ins_err)
		}
	} else {
		rows, err := db.Query("select name,age,id,carrier,imageurl,snippet from phones")
		PanicOn(err)

		for rows.Next(){
			var name string
			var age int
			var id string
			var carrier string
			var imageurl string
			var snippet string
			err = rows.Scan(&name,&age,&id,&carrier,&imageurl,&snippet)
			allPhones = append(allPhones, Phone{name,age,id,carrier,imageurl,snippet})
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
