package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomer(t *testing.T) {
	assert := assert.New(t)
	assert.True(true)
}

func init() {
	connection, err := sqlx.Open("postgres", "dbname=test user=postgres password=postgres sslmode=disable")
	PanicOn(err)
	db = connection

	clearTables()
}

func TestNewCustomer(t *testing.T) {
	assert := assert.New(t)

	c, err := NewCustomer(db, "testNewCustomer@gmail.com", "new", "test", time.Now())
	assert.Nil(err)
	assert.Equal("testNewCustomer@gmail.com", c.Email)
	assert.Equal("new", c.FirstName)
	assert.Equal("test", c.LastName)
}

func TestDeleteCustomer(t *testing.T) {
	assert := assert.New(t)
	var count int

	newC, err := NewCustomer(db, "testDeleteCustomer@gmail.com", "delete", "test", time.Now())
	assert.Nil(err)

	query := `SELECT COUNT(*) FROM customers`
	err = db.Get(&count, query)
	assert.Nil(err)
	before := count

	err = DeleteCustomer(db, newC.ID)
	assert.Nil(err)

	err = db.Get(&count, query)
	assert.Nil(err)
	assert.Equal(before-1, count)
}

func TestUpdateCustomer(t *testing.T) {
	assert := assert.New(t)
	var c Customer

	newC, err := NewCustomer(db, "testUpdateCustomer@gmail.com", "update", "test", time.Now())
	assert.Nil(err)

	newC.Email = "changed@gmail.com"
	newC.FirstName = "changed"
	newC.LastName = "name"

	err = UpdateCustomer(db, newC)
	assert.Nil(err)

	query := `SELECT * FROM customers WHERE customer_id = ($1)`
	err = db.Get(&c, query, newC.ID)

	assert.Equal(newC.Email, c.Email)
	assert.Equal(newC.FirstName, c.FirstName)
	assert.Equal(newC.LastName, c.LastName)
}

func TestFindCustomerByEmail(t *testing.T) {
	assert := assert.New(t)

	newC, err := NewCustomer(db, "testFindCustomerEmail@gmail.com", "find", "email", time.Now())
	assert.Nil(err)

	c, err := FindCustomerByEmail(db, newC.Email)
	assert.Equal(newC.Email, c.Email)
	assert.Equal(newC.FirstName, c.FirstName)
	assert.Equal(newC.LastName, c.LastName)
}

func TestFindCustomerByID(t *testing.T) {
	assert := assert.New(t)

	newC, err := NewCustomer(db, "testFindCustomerID@gmail.com", "find", "ID", time.Now())
	assert.Nil(err)

	c, err := FindCustomerByID(db, newC.ID)
	assert.Equal(newC.Email, c.Email)
	assert.Equal(newC.FirstName, c.FirstName)
	assert.Equal(newC.LastName, c.LastName)
}

func TestAllCustomers(t *testing.T) {
	assert := assert.New(t)
	var count int

	_, err := NewCustomer(db, "testAllCustomers@gmail.com", "all", "customers", time.Now())
	assert.Nil(err)

	query := `SELECT COUNT(*) FROM customers`
	err = db.Get(&count, query)
	assert.Nil(err)

	cArray, err := AllCustomers(db)
	assert.Equal(count, len(cArray))
}
