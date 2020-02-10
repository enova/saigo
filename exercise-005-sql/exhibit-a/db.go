package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}
