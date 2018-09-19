package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Person struct {
	PersonId int    `db:"person_id"`
	Name     string `db:"name"`
	Ssn      int    `db:"ssn"`
}

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var _ sql.Result
	db, err := sqlx.Open("postgres", "dbname=exercise-005-sql sslmode=disable")
	PanicOn(err)
	cleanDB(db)
	seedData(db)
	mutateData(db)

}

const TO_BE_DELETED_SSN = 111223333
const TO_BE_UPDATED_SSN_FROM = 444556666
const TO_BE_UPDATED_SSN_TO = 444556668
const TO_BE_INSERTED_SSN = 444556667

func cleanDB(db *sqlx.DB) {
	db.Exec("TRUNCATE people")
}
func seedData(db *sqlx.DB) {

	_, err := db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", TO_BE_DELETED_SSN)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", TO_BE_UPDATED_SSN_FROM)
	PanicOn(err)

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)
	var id int
	var name string
	var ssn int
	for rows.Next() {

		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}

func mutateData(db *sqlx.DB) {
	var count int
	var person Person

	fmt.Println("1. Delete a person")
	_, err := db.Exec("DELETE FROM  people WHERE ssn = $1", TO_BE_INSERTED_SSN)
	PanicOn(err)

	err = db.Get(&count, "SELECT COUNT(1) AS total FROM people WHERE ssn = $1", TO_BE_UPDATED_SSN_TO)
	PanicOn(err)
	if count != 0 {
		PanicOn(errors.New("Failed to delete"))
	}

	fmt.Println("2. Update a person's SSN")
	_, err = db.Exec("UPDATE people SET ssn = $1 WHERE ssn = $2 ",
		TO_BE_UPDATED_SSN_TO,
		TO_BE_UPDATED_SSN_FROM)
	PanicOn(err)

	err = db.Get(&count, "SELECT COUNT(1) AS total FROM people WHERE ssn = $1", TO_BE_UPDATED_SSN_TO)
	PanicOn(err)
	if count != 1 {
		PanicOn(errors.New("Failed to update"))
	}

	fmt.Println("3. Insert a person and immediately retrieve the person's ID assigned by the database")

	err = db.Get(&person, `INSERT INTO people (name, ssn) VALUES ($1, $2) RETURNING *`,
		"Smith",
		TO_BE_INSERTED_SSN)
	PanicOn(err)
	if person.Ssn != TO_BE_INSERTED_SSN {
		PanicOn(errors.New("Failed to insert"))
	}

	fmt.Println("Tests passed: ok")
}
