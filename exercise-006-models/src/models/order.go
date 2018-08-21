package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// Order ...
type Order struct {
	ID         int       `db:"order_id"`
	ProductID  int       `db:"product_id"`
	Quantity   int       `db:"quantity"`
	CustomerID int       `db:"customer_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// NewOrder ...
func NewOrder(db *sqlx.DB, customerID int, productID int, quantity int) (*Order, error) {
	var o Order
	query := `INSERT INTO orders(customer_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING *`
	err := db.Get(&o, query, customerID, productID, quantity)
	return &o, err
}

// UpdateOrder ...
func UpdateOrder(db *sqlx.DB, o *Order) error {
	query := `UPDATE orders SET product_id = ($1), quantity = ($2), customer_id = ($3), created_at = ($4), updated_at = ($5) WHERE order_id = ($6)`
	_, err := db.Exec(query, o.ProductID, o.Quantity, o.CustomerID, o.CreatedAt, o.UpdatedAt, o.ID)
	return err
}

// DeleteOrder ...
func DeleteOrder(db *sqlx.DB, orderID int) error {
	query := `DELETE FROM orders WHERE order_id = ($1)`
	_, err := db.Exec(query, orderID)
	return err
}
