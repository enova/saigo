package main

import (
	"database/sql"
)

// People is a singleton for communicating with the people table
var People peopleConn

type peopleConn struct {
	Db *sql.DB
}

func (p peopleConn) DeleteAll() error {
	_, err := p.Db.Exec("DELETE FROM people")
	return err
}

func (p peopleConn) Delete(id int64) error {
	_, err := p.Db.Exec("DELETE FROM people WHERE person_id=$1", id)
	return err
}

func (p peopleConn) Update(id int64, ssn int) error {
	_, err := p.Db.Exec("UPDATE people SET ssn=$1 WHERE person_id=$2", ssn, id)
	return err
}

func (p peopleConn) Insert(name string, ssn int) (int64, error) {
	var userid int64
	row := p.Db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1, $2) RETURNING person_id", name, ssn)
	err := row.Scan(&userid)
	return userid, err
}

func (p peopleConn) AllRows() (*sql.Rows, error) {
	return p.Db.Query("SELECT * FROM people")
}

func (p peopleConn) All() []Person {
	people := make([]Person, 0)

	rows, err := p.AllRows()
	if err != nil {
		return people
	}

	for rows.Next() {
		p, err := Person{}.InitFromRow(rows)
		if err == nil {
			people = append(people, p)
		}
	}
	return people
}

// OpenPeopleConnection opens a connection to the database
func OpenPeopleConnection() {
	People = peopleConn{}
	People.Db, _ = sql.Open("postgres", "dbname=test sslmode=disable")
}

// ClosePeopleConnection closes the open connection to the database
func ClosePeopleConnection() {
	People.Db.Close()
}
