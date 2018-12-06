
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

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

	//printPeople(db)

	// Delete person by name "Sho 'Nuff" from people table
	_, err = db.Exec("DELETE FROM people where name=$1", "Sho 'Nuff")
	PanicOn(err)

	// Update ssn of a person in people table
	_, err = db.Exec("UPDATE people SET ssn=$1 where name=$2", 12121323, "Bruce Leroy")
	PanicOn(err)

	// Insert a new person and retrieve the ID
	var id int
	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "Akash", "121213242").Scan(&id)
	PanicOn(err)

	fmt.Println("Id of the person inserted is ", id)

	printPeople(db)

}

func printPeople(db *sql.DB){

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