package models

import (
	"database/sql"
)

// Product ...
type Product struct {
	ID      int
	Product string
}

// NewProduct ...
func NewProduct(db *sql.DB, productName string) error {
	return nil
}

// UpdateProduct ...
func UpdateProduct(db *sql.DB, o *Product) error {
	return nil
}

// DeleteProduct ...
func DeleteProduct(db *sql.DB, productID int) error {
	return nil
}

// FindProduct ...
func FindProduct(db *sql.DB, productName string) (*Product, error) {
	return nil, nil
}
