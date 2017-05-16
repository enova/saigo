## Description
Learn to work with SQL in Go!

## Prerequisites
This exercise assumes the learner is familiar with basic SQL. If you are new to SQL then you must first
learn the basics before proceeding. Saigo uses [PostgreSQL](http://www.tutorialspoint.com/postgresql/) for all database work.

Set up a local PostgreSQL server on your deveopment machine and create a database called `test`. Connect to the database
(`$ psql test`) and copy the code in `migration_up.sql` to create the `people` table in your `test` database:

```
CREATE TABLE people (
    person_id SERIAL PRIMARY KEY,
    name TEXT,
    ssn INTEGER
);
```

## Comprehension Tasks

The standard Go distribution includes a `database/sql` package with implementations for various databases, including PostgreSQL.
A nice extension to `database/sql` is the [`sqlx`](https://github.com/jmoiron/sqlx) package that uses some reflection to create a slightly
prettier interface.

Check out the code in [exhibit-a]():

```go
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
	defer db.Close()

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

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
```

Make sure you followed the steps above to create the `people` table in your `test` database. Run this code on the console
and see it work. If you have issues seek the assistance of an instructor.

Explain what each line of code does in `exhibit-a`.

# Engineering Tasks

Create an `exhibit-b` folder and expand the code (from `exhibit-a`) to:

1. Delete a person
1. Update a person's SSN
1. Insert a person and immediately retrieve the person's ID assigned by the database
1. Handle errors that might be returned by each of the function calls made on `db`
