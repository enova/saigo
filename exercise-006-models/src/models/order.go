package models

import (
	"database/sql"
	"time"
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
func NewOrder(db *sql.DB, customerID int, productID int, quantity int) error {
	return nil
}

// UpdateOrder ...
func UpdateOrder(db *sql.DB, o *Order) error {
	return nil
}

// DeleteOrder ...
func DeleteOrder(db *sql.DB, orderID int) error {
	return nil
}
