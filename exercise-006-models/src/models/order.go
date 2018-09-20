package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Order ...
type Order struct {
	ID         int `db:"order_id"`
	ProductID  int `db:"product_id"`
	Quantity   int
	CustomerID int       `db:"customer_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// NewOrder ...
func NewOrder(db *sqlx.DB, customerID int, productID int, quantity int) error {
	_, err := db.Exec("INSERT INTO orders(customer_id, product_id, quantity) values($1, $2, $3)", customerID, productID, quantity)
	return err
}

// Update Order ...
func UpdateOrder(db *sqlx.DB, o *Order) error {
	_, err := db.NamedExec("UPDATE orders SET customer_id=:customer_id, product_id=:product_id, quantity=:quantity, updated_at=:updated_at WHERE order_id=:order_id", o)
	return err
}

// DeleteOrder ...
func DeleteOrder(db *sqlx.DB, orderID int) error {
	_, err := db.Exec("DELETE FROM orders where order_id=$1", orderID)
	return err
}
