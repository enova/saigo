package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"io/ioutil"
	"net/http"
)

var Conn *sqlx.DB


const DB_URI = "user=postgres dbname=exercise-007-json sslmode=disable"


func GetConn() *sqlx.DB {
	if Conn == nil{
		var err error
		Conn, err = sqlx.Connect("postgres", DB_URI)
		if err != nil {
			fmt.Println(err)
		}

	}
	return Conn
}

// Phone ...
type Phone struct {
	ID string `json:"id" db:"id"`
	Age int `json:"age" db:"age"`
	ImageUrl string `json:"imageUrl" db:"image_url"`
	Name string `json:"name" db:"name"`
	Snippet string `json:"snippet" db:"snippet"`
}

func setup() {
	var count int
	err := GetConn().Get(&count, "SELECT COUNT(1) FROM phones")
	if err!= nil{
		panic(err)
	}
	if(count == 0){
		cpJsonToDb()
	}

}
func cpJsonToDb() {
	data, err := ioutil.ReadFile("phones.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var phones []Phone
	err = json.Unmarshal(data, &phones)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
		panic(err)
	}

	tx, err := GetConn().Beginx()
	for _, value := range phones {
		query := `
			INSERT INTO phones (id, age, image_url, name, snippet) 
				VALUES(:id, :age, :image_url, :name, :snippet)`
		tx.NamedExec(query, &value)
	}
	err = tx.Commit()
	if err != nil{
		panic(err)
	}
}

func handlePhones(w http.ResponseWriter, r *http.Request) {
	var phones []Phone
	GetConn().Select(&phones, "SELECT * FROM phones ORDER BY id ASC")
	json.NewEncoder(w).Encode(phones)
}

func main() {
	setup()
	http.HandleFunc("/phones", handlePhones)
	http.ListenAndServe(":8080", nil)
}
