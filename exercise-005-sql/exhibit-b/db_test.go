package main

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var db = new(sql.DB)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("postgres", "dbname=test sslmode=disable")
	if err != nil {
		fmt.Println("Error opening database connection")
		os.Exit(1)
	}
	code := m.Run()
	os.Exit(code)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	tx, _ := db.Begin()
	testID := 1
	row := get(tx, testID)
	var id int
	var name string
	var ssn int
	row.Scan(&id, &name, &ssn)
	assert.Equal(testID, id)
	tx.Rollback()
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	tx, _ := db.Begin()
	i, err := insert(tx, "testing", 123121234)
	assert.Nil(err)
	assert.NotZero(i)
	tx.Rollback()
}

func TestUpdate(t *testing.T) {
	testID := 1
	testSSN := 111111111
	assert := assert.New(t)
	tx, _ := db.Begin()
	err := update(tx, testID, testSSN)
	assert.Nil(err)
	row := get(tx, testID)
	var id int
	var name string
	var ssn int
	row.Scan(&id, &name, &ssn)
	assert.Equal(testID, id)
	assert.Equal(testSSN, ssn)
	tx.Rollback()
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	tx, _ := db.Begin()
	err := delete(tx, 1)
	assert.Nil(err)
	tx.Rollback()
}
