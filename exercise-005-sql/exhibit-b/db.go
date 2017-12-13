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
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)
	defer db.Close()

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

	_, err = db.Exec("DELETE FROM people WHERE ssn = $1", 111223333)
	PanicOn(err)

	_, err = db.Exec("UPDATE people SET ssn = $1 WHERE ssn = $2", 999887777, 444556666)
	PanicOn(err)

	res, err := db.Query("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "Jim Bob", 1234567789)
	PanicOn(err)
	defer res.Close()

	var id int
	res.Next()
	res.Scan(&id)
	fmt.Printf("Last Inserted id: %d\n", id)
	res.Close()

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}
