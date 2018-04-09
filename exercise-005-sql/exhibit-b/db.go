package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"github.com/saigo/exercise-005-sql/exhibit-b/utils"
)

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

// Task 4: Handle errors that might be returned by each of the function calls made on db
//
func CheckIfError(err error) bool {
	if nil == err {
		return false
	}

	fmt.Println(utils.WhereAmI(2), "- error:", err)
	return true
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	if (CheckIfError(err)) { return }

	// Clear the table
	//
	_, err = db.Query("DELETE FROM people")
	if (CheckIfError(err)) { return }

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	if (CheckIfError(err)) { return }

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	if (CheckIfError(err)) { return }

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	if (CheckIfError(err)) { return }

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}

	// Task 1: Delete person
	//
	name := "Bruce Leroy"
	_, err = db.Query("DELETE FROM people where name='" + name + "'")
	if (CheckIfError(err)) { return }

	rows, err = db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}

	// Task 2: Update person's SSN
	//
	name = "Sho 'Nuff"
	ssn := 123456789
	_, err = db.Query("UPDATE people set ssn=$1 WHERE name=$2", strconv.Itoa(ssn), name)
	if (CheckIfError(err)) { return }

	rows, err = db.Query("SELECT person_id, name, ssn FROM people")
	if (CheckIfError(err)) { return }

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}

	// Task 4: Insert a person and immediately retrieve the person's ID assigned by the database
	//
	name = "Jay Leno"
	ssn = 987654321
	var id int
	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", name, ssn).Scan(&id)
	if (CheckIfError(err)) { return }
	fmt.Printf("Inserted person %5d %-15s %9d\n", id, name, ssn)

	rows, err = db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)
	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		if (CheckIfError(err)) { return }

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}

}
