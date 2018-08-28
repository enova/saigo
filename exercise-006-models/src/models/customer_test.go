package models

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func init() {
	conn, err := sqlx.Connect("postgres", "dbname=test sslmode=disable")
	panicOn(err)
	db = conn
}

func TestCustomerRefresh(t *testing.T) {

}

func TestNewCustomer(t *testing.T) {
	// Create a customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	// Make sure the customer was created in the DB
	var count int
	db.QueryRowx(`SELECT count(*) FROM customers WHERE customer_id=$1`, customer.ID).Scan(&count)
	assert.Equal(t, count, 1)
}

func TestDeleteCustomer(t *testing.T) {
	// Create a customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)

	// Delete the customer
	DeleteCustomer(db, customer.ID)

	// Make sure the customer does not exist in the DB
	var count int
	db.QueryRowx(`SELECT count(*) FROM customers WHERE customer_id=$1`, customer.ID).Scan(&count)
	assert.Equal(t, count, 0)
}

func TestUpdateCustomer(t *testing.T) {
	// Create a customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	// Update the customer first name
	updatedFirstName := "newName"
	udpatedCustomer := Customer{ID: customer.ID, Email: customer.Email, FirstName: updatedFirstName, LastName: customer.LastName, BirthDate: customer.BirthDate}
	err = UpdateCustomer(db, &udpatedCustomer)
	assert.Nil(t, err)

	// make sure the customer first name is updated
	retrievedCustomer, err := FindCustomerByID(db, customer.ID)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedCustomer)
	assert.Equal(t, retrievedCustomer.FirstName, updatedFirstName)

}

func TestFindCustomerByEmail(t *testing.T) {
	// Create a customer
	email := "test@example.com"
	customer, err := NewCustomer(db, email, "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	retrievedCustomer, err := FindCustomerByEmail(db, email)
	assert.Nil(t, err)
	assert.Equal(t, retrievedCustomer.ID, customer.ID)
}

func TestFindCustomerByID(t *testing.T) {
	// Create a customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	retrievedCustomer, err := FindCustomerByID(db, customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, retrievedCustomer.Email, customer.Email)
}

func TestAllCustomers(t *testing.T) {
	// Create a customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	// Create another customer
	customer2, err := NewCustomer(db, "test2@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer2.ID)

	customers, err := AllCustomers(db)
	assert.Nil(t, err)
	assert.Equal(t, len(customers), 2)
}
