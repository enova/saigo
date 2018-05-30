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

  _, err = db.Exec("DROP table people")
  PanicOn(err)

  _, err = db.Exec("CREATE TABLE people (person_id SERIAL PRIMARY KEY, name TEXT, ssn INTEGER)")
  PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

  _, err = db.Exec("Update people set ssn=$1 where name=$2", 123456789, "Sho 'Nuff")
  PanicOn(err)

  _, err = db.Exec("Delete from people where name = $1", "Bruce Leroy")
  PanicOn(err)

  val, err := db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 777889999)
  PanicOn(err)
  id, err := val.LastInsertId()
  fmt.Printf("%d\n", id)

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
