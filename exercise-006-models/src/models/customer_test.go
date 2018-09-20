package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	conn, err := sqlx.Connect("postgres", "dbname=test sslmode=disable")
	PanicOn(err)
	db = conn
}

func TestNewCustomer(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	var count int
	db.QueryRowx("SELECT count(*) FROM customers WHERE customer_id=$1", customer.ID).Scan(&count)
	assert.Equal(t, count, 1)
}

func TestDeleteCustomer(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)

	DeleteCustomer(db, customer.ID)

	var count int
	db.QueryRowx("SELECT count(*) FROM customers WHERE customer_id=$1", customer.ID).Scan(&count)
	assert.Equal(t, count, 0)
}

func TestUpdateCustomer(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	newName := "Sho"
	updatedCustomer := Customer{ID: customer.ID, Email: customer.Email, FirstName: newName, LastName: customer.LastName, BirthDate: customer.BirthDate}
	err = UpdateCustomer(db, &updatedCustomer)
	assert.Nil(t, err)

	actualCustomer, err := FindCustomerByID(db, customer.ID)
	assert.Nil(t, err)
	assert.NotNil(t, actualCustomer)
	assert.Equal(t, actualCustomer.FirstName, newName)
}

func TestFindCustomerByEmail(t *testing.T) {
	email := "cust1@test.com"
	customer, err := NewCustomer(db, email, "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	actualCustomer, err := FindCustomerByEmail(db, email)
	assert.Nil(t, err)
	assert.Equal(t, actualCustomer.ID, customer.ID)
}

func TestFindCustomerByID(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	actualCustomer, err := FindCustomerByID(db, customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, actualCustomer.Email, customer.Email)
}

func TestAllCustomers(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	customer2, err := NewCustomer(db, "cust2@test.com", "Sho", "Nuff", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer2.ID)

	customers, err := AllCustomers(db)
	assert.Nil(t, err)
	assert.Equal(t, len(customers), 2)
}
