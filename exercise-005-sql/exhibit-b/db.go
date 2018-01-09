package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// LogError will log.Fatal error if error is given
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	LogError(err)
	defer db.Close()

	// For sake of this exercise, clear the DB first to ensure everything runs ok following this
	_, err = db.Exec("DELETE FROM people;")
	LogError(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2);", "Bruce Leroy", 111223333)
	LogError(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2);", "Sho testtest", 444556655)
	LogError(err)

	id := -1
	// since err could be set from the call above, set it to nil since we check it below
	err = nil
	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id;", "test bob", 442256516).Scan(&id)
	LogError(err)
	if err == nil && id != -1 {
		_, err = db.Exec("DELETE FROM people WHERE person_id = $1;", id)
		LogError(err)
	}

	_, err = db.Exec("UPDATE people SET ssn = $1 WHERE name = $2;", 123675874, "Bruce Leroy")
	LogError(err)
}
