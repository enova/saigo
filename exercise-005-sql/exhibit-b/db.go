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

	_, err = db.Exec("DELETE from people where name = 'Bruce Leroy'")
	PanicOn(err)

	_, err = db.Exec("UPDATE people SET ssn = 111111111 where name = 'Sho ''Nuff'")
	PanicOn(err)

	var pid int

	err = db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", "Bruce Leroy", 222222222).Scan(&pid)
	PanicOn(err)

	fmt.Printf("Person %5d\n", pid)

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
