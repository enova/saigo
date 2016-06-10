package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Order ...
type Order struct {
	ID         int
	ProductID  int
	Quantity   int
	CustomerID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewOrder ...
func NewOrder(db *sqlx.DB, customerID int, productID int, quantity int) error {
	return nil
}

// Update Order ...
func UpdateOrder(db *sqlx.DB, o *Order) error {
	return nil
}

// DeleteOrder ...
func DeleteOrder(db *sqlx.DB, orderID int) error {
	return nil
}
