package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// PanicOn will panic if error is given
func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho testtest", 444556655)
	PanicOn(err)

	var id int
	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "test bob", 442256516).Scan(&id)
	PanicOn(err)

	_, err = db.Exec("DELETE FROM people WHERE person_id = $1;", id)
	PanicOn(err)

	_, err = db.Exec("UPDATE people set ssn = $1 WHERE name = $2", 123675874, "Bruce Leroy")
	PanicOn(err)
}
