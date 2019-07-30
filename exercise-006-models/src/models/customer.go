package models

import (
	"time"

	"github.com/jmoiron/sqlx"
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
func (c *Customer) Refresh(db *sqlx.DB) error {
	return nil
}

// NewCustomer ...
func NewCustomer(db *sqlx.DB, email string, firstName string, lastName string, birthDate time.Time) (*Customer, error) {
	return nil, nil
}

// DeleteCustomer ...
func DeleteCustomer(db *sqlx.DB, id int) error {
	return nil
}

// UpdateCustomer ...
func UpdateCustomer(db *sqlx.DB, u *Customer) error {
	return nil
}

// FindCustomerByEmail ...
func FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error) {
	return nil, nil
}

// FindCustomerByID ...
func FindCustomerByID(db *sqlx.DB, id int) (*Customer, error) {
	return nil, nil
}

// AllCustomers ...
func AllCustomers(db *sqlx.DB) ([]*Customer, error) {
	return nil, nil
}
