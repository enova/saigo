package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Customer ...
type Customer struct {
	ID        int `db:"customer_id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	BirthDate time.Time `db:"birth_date"`
	Orders    []* Order
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Refresh ...
func (c *Customer) Refresh(db *sqlx.DB) error {
	return db.Get(c, "SELECT * FROM customers WHERE customer_id = $1", c.ID)
}

// NewCustomer ...
func NewCustomer(db *sqlx.DB, email string, first_name string, last_name string, birth_date time.Time) (*Customer, error) {
	var u Customer
	query := "INSERT INTO customers (email, first_name, last_name, birth_date) VALUES($1,$2,$3,$4) RETURNING *"
	err := db.Get(&u, query, email, first_name, last_name, birth_date)
	return &u, err
}

// DeleteCustomer ...
func DeleteCustomer(db *sqlx.DB, id int) error {
	query := "DELETE FROM customers WHERE customer_id = $1"
	_, err := db.Exec(query, id)
	return err
}

// UpdateCustomer ...
func UpdateCustomer(db *sqlx.DB, u *Customer) error {
	query := `
		UPDATE customers
			SET
				email = :email,
				first_name = :first_name,
				last_name = :last_name,
				birth_date = :birth_date,
				updated_at = NOW()
			WHERE customer_id = :customer_id`
	_, err := db.NamedExec(query, u)
	return err
}

// FindCustomerByEmail ...
func FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error) {
	var u Customer
	err := db.Get(&u, "SELECT * FROM customers WHERE email = $1", email)
	return &u, err
}

// FindCustomerByID ...
func FindCustomerByID(db *sqlx.DB, id int) (*Customer, error) {
	var u Customer
	err := db.Get(&u, "SELECT * FROM customers WHERE customer_id = $1", id)
	return &u, err
}

// AllCustomers ...
func AllCustomers(db *sqlx.DB) ([]*Customer, error) {
	var users []*Customer
	err := db.Select(&users, "SELECT * FROM customers")
	return users, err
}
