package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	panicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES($1, $2)", "Bruce Leroy", 111223333)
	panicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES($1, $2)", "Sho 'Nuff", 444556666)
	panicOn(err)

	printAllPeople(db)

	fmt.Println("-------------------------------------")

	// Delete a person
	_, err = db.Exec("DELETE FROM people where name=$1", "Bruce Leroy")
	panicOn(err)

	// Update a person's SSN
	_, err = db.Exec("UPDATE people SET ssn=$1 where name=$2", 9999, "Sho 'Nuff")
	panicOn(err)

	// Insert a new person and immidiately retrieve the person's ID assigned by the database
	var id int
	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "akash", 1234).Scan(&id)
	panicOn(err)
	fmt.Println("Id for new person:", id)

	printAllPeople(db)
}

func printAllPeople(db *sql.DB) {
	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	panicOn(err)

	for rows.Next() {
		var (
			id   int
			name string
			ssn  int
		)

		err := rows.Scan(&id, &name, &ssn)
		panicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}
