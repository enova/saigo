package main

import (
	"database/sql"
	"fmt"
)

// Person is an ORM for the people table.
type Person struct {
	ID   int
	Name string
	Ssn  int
}

// ToString returns a string representation of this Person
func (p Person) ToString() string {
	return fmt.Sprintf("%5d %-15s %9d\n", p.ID, p.Name, p.Ssn)
}

// InitFromRow initializes a person from a row of sql data
func (p Person) InitFromRow(row *sql.Rows) (Person, error) {
	err := row.Scan(&p.ID, &p.Name, &p.Ssn)
	return p, err
}
