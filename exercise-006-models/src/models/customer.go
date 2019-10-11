package models

import (
	"database/sql"
	"time"
)

// Customer ...
type Customer struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	BirthDate time.Time
	Orders    []*Order
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Refresh ...
func (c *Customer) Refresh(db *sql.DB) error {
	return nil
}

// NewCustomer ...
func NewCustomer(db *sql.DB, email string, firstName string, lastName string, birthDate time.Time) (*Customer, error) {
	return nil, nil
}

// DeleteCustomer ...
func DeleteCustomer(db *sql.DB, id int) error {
	return nil
}

// UpdateCustomer ...
func UpdateCustomer(db *sql.DB, u *Customer) error {
	return nil
}

// FindCustomerByEmail ...
func FindCustomerByEmail(db *sql.DB, email string) (*Customer, error) {
	return nil, nil
}

// FindCustomerByID ...
func FindCustomerByID(db *sql.DB, id int) (*Customer, error) {
	return nil, nil
}

// AllCustomers ...
func AllCustomers(db *sql.DB) ([]*Customer, error) {
	return nil, nil
}
