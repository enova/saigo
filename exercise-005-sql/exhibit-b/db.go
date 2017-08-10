package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func panicOn(err error, db *sql.DB) {
	if err != nil {
		fmt.Println(err)
		db.Close()
		os.Exit(1)
	}
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	panicOn(err, db)

	_, err = db.Exec("DROP TABLE IF EXISTS people")
	panicOn(err, db)

	_, err = db.Exec("CREATE TABLE people (person_id SERIAL PRIMARY KEY, name TEXT, ssn INTEGER);")
	panicOn(err, db)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	panicOn(err, db)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	panicOn(err, db)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Asdf Fdsa", 123456789)
	panicOn(err, db)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Zxcv Vcxz", 987654321)
	panicOn(err, db)

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	panicOn(err, db)

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		panicOn(err, db)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}

	_, err = db.Exec("DELETE FROM people WHERE person_id > 2")
	panicOn(err, db)

	_, err = db.Exec("UPDATE people SET name = 'Asdf Zxcv' WHERE person_id = 3")
	panicOn(err, db)

	row := db.QueryRow("INSERT INTO people(name, ssn) VALUES($1, $2) RETURNING person_id", "Zxcv Asdf", 000000000)

	var id int
	err = row.Scan(&id)
	panicOn(err, db)

	fmt.Printf("Person %5d\n", id)

	_, err = db.Exec("DROP TABLE people")
	panicOn(err, db)

	err = db.Close()
	panicOn(err, db)
}
