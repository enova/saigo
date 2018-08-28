package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Customer ...
type Customer struct {
	ID        int `db:"customer_id"`
	Email     string
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	BirthDate time.Time `db:"birth_date"`
	Orders    []*Order
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Refresh ...
func (c *Customer) Refresh(db *sqlx.DB) error {
	err := db.Get(c, `SELECT * FROM customers WHERE customer_id=$1`, c.ID)
	return err
}

// NewCustomer ...
func NewCustomer(db *sqlx.DB, email string, first_name string, last_name string, birth_date time.Time) (*Customer, error) {
	var customerID int
	err := db.QueryRowx(`INSERT INTO customers(email, first_name, last_name, birth_date) values ($1, $2, $3, $4) returning customer_id`, email, first_name, last_name, birth_date).Scan(&customerID)
	if err != nil {
		return nil, err
	}

	customer := &Customer{}
	err = db.Get(customer, `SELECT * FROM customers WHERE customer_id=$1`, customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// DeleteCustomer ...
func DeleteCustomer(db *sqlx.DB, id int) error {
	_, err := db.Exec(`DELETE FROM customers WHERE customer_id=$1`, id)
	return err
}

// UpdateCustomer ...
func UpdateCustomer(db *sqlx.DB, u *Customer) error {
	_, err := db.NamedExec(`UPDATE customers SET email=:email, first_name=:first_name, last_name=:last_name, birth_date=:birth_date, updated_at=:updated_at WHERE customer_id=:customer_id`, u)
	return err
}

// FindCustomerByEmail ...
func FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error) {
	customer := &Customer{}
	err := db.Get(customer, `SELECT * FROM customers WHERE email=$1`, email)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// FindCustomerByID ...
func FindCustomerByID(db *sqlx.DB, id int) (*Customer, error) {
	customer := &Customer{}
	err := db.Get(customer, `SELECT * FROM customers WHERE customer_id=$1`, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// AllCustomers ...
func AllCustomers(db *sqlx.DB) ([]*Customer, error) {
	customers := []*Customer{}
	rows, err := db.Queryx(`SELECT * FROM customers`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		customer := &Customer{}
		rows.StructScan(customer)
		customers = append(customers, customer)
	}

	return customers, nil
}
