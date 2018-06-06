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


	_, err = db.Exec("DELETE FROM people where name = $1", "Bruce Leroy")
	PanicOn(err)

	_, err = db.Exec("UPDATE PEOPLE set ssn=$1 WHERE NAME=$2", 999999999, "Sho 'Nuff")
	PanicOn(err)

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "NICK", 123456789)
	PanicOn(err)

	rowLastInsert, err := db.Query("SELECT person_id FROM people WHERE name=$1", "NICK")
	PanicOn(err)

	for rowLastInsert.Next() {
		var id int
		err := rowLastInsert.Scan(&id)
		PanicOn(err)

		fmt.Printf("Last Insert Person %5d\n", id)
	}

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}
