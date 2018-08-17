package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)

	// Add 2 entries
	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

	// Delete a person
	_, err = db.Exec("DELETE FROM people WHERE name=($1)", "Bruce Leroy")
	PanicOn(err)

	// Update a person's SSN
	_, err = db.Exec("UPDATE people SET ssn = ($1) WHERE person_id=($2)", 124151236, 2)
	PanicOn(err)

	// Insert a person and immediately retrieve the person's ID assigned by the database
	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "Kevin Tan", 123233456)
	PanicOn(err)

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}
