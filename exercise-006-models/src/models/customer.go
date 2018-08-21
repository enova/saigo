package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// Customer user information
type Customer struct {
	ID        int       `db:"customer_id"`
	Email     string    `db:"email"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	BirthDate time.Time `db:"birth_date"`
	Orders    []*Order
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Refresh ...update the values of the receiving Customer instance with the latest values from the database
func (c *Customer) Refresh(db *sqlx.DB) error {
	query := `SELECT * FROM customers WHERE customer_id = ($1)`
	err := db.Get(c, query, c.ID)
	return err
}

// NewCustomer ...
func NewCustomer(db *sqlx.DB, email string, firstName string, lastName string, birthDate time.Time) (*Customer, error) {
	var c Customer
	query := `INSERT INTO customers(email, first_name, last_name, birth_date) VALUES ($1, $2, $3, $4) RETURNING *`
	err := db.Get(&c, query, email, firstName, lastName, birthDate)
	return &c, err
}

// DeleteCustomer ...
func DeleteCustomer(db *sqlx.DB, id int) error {
	query := `DELETE FROM customers WHERE customer_id = ($1)`
	_, err := db.Exec(query, id)
	return err
}

// UpdateCustomer ...
func UpdateCustomer(db *sqlx.DB, u *Customer) error {
	query := `UPDATE customers SET email = ($1), first_name = ($2), last_name = ($3), birth_date = ($4), created_at = ($5), updated_at = ($6) WHERE customer_id = ($7)`
	_, err := db.Exec(query, u.Email, u.FirstName, u.LastName, u.BirthDate, u.CreatedAt, u.UpdatedAt, u.ID)
	return err
}

// FindCustomerByEmail ...
func FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error) {
	var c Customer
	query := `SELECT * FROM customers WHERE email = ($1)`
	err := db.Get(&c, query, email)
	return &c, err
}

// FindCustomerByID ...
func FindCustomerByID(db *sqlx.DB, id int) (*Customer, error) {
	var c Customer
	query := `SELECT * FROM customers WHERE customer_id = ($1)`
	err := db.Get(&c, query, id)
	return &c, err
}

// AllCustomers ...
func AllCustomers(db *sqlx.DB) ([]*Customer, error) {
	var c []*Customer
	query := `SELECT * FROM customers`
	err := db.Select(&c, query)
	return c, err
}
