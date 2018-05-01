package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	if err != nil {
		fmt.Printf("Error while opening DB: %s\n\n", err)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Error starting transaction: %s\n\n", err)
		return
	}

	show(tx)

	id, err := insert(tx, "Dave", 111001111)
	if err != nil {
		fmt.Printf("Error on insert: %s\n\n", err)
		return
	}

	show(tx)

	err = update(tx, id, 111221111)
	if err != nil {
		fmt.Printf("Error on update: %s\n\n", err)
		return
	}

	show(tx)

	err = delete(tx, id)
	if err != nil {
		fmt.Printf("Error on delete %s\n\n", err)
		return
	}

	show(tx)

	tx.Commit()
}

func show(tx *sql.Tx) error {
	rows, err := tx.Query("SELECT person_id, name, ssn FROM people")
	if err != nil {
		fmt.Printf("Error getting people from DB: %s\n\n", err)
		return err
	}

	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		if err != nil {
			fmt.Printf("Error getting row: %s\n\n", err)
			return err
		}

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
	fmt.Println()
	return nil
}

func get(tx *sql.Tx, id int) *sql.Row {
	row := tx.QueryRow("SELECT person_id, name, ssn FROM people WHERE person_id=$1", id)
	return row
}

func update(tx *sql.Tx, id int, ssn int) error {
	_, err := tx.Exec("UPDATE people SET ssn=$1 WHERE person_id=$2", ssn, id)
	return err
}

func insert(tx *sql.Tx, name string, ssn int) (int, error) {
	key := 0
	err := tx.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", name, ssn).Scan(&key)
	if err != nil {
		return 0, err
	}
	return int(key), nil
}

func delete(tx *sql.Tx, id int) error {
	_, err := tx.Exec("DELETE FROM people WHERE person_id = $1", id)
	return err
}
