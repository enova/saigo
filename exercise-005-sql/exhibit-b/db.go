package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func printAllPeople(db *sql.DB)  {
	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)

	fmt.Println("\n------------------------------------------------")

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}

	fmt.Println("------------------------------------------------\n")
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)

	// Let's truncate table
	_, err = db.Exec("TRUNCATE TABLE people")
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

	printAllPeople(db)

	// delete Bruce
	fmt.Println("Deleting Bruce Leroy")
	_, err = db.Exec("DELETE FROM people WHERE name = 'Bruce Leroy'")
	PanicOn(err)

	fmt.Println("Updating ssn")
	_, err = db.Exec("UPDATE people SET ssn = 5443555 WHERE name = $1", "Sho 'Nuff")
	PanicOn(err)

	printAllPeople(db)

	var id int

	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "A Max", 444556667).Scan(&id)

	PanicOn(err)
	fmt.Printf("Person %5d\n", id)

	printAllPeople(db)

}
