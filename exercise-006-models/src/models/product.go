package models

import (
	"github.com/jmoiron/sqlx"
)

// Product ...
type Product struct {
	ID      int
	Product string
}

// NewProduct ...
func NewProduct(db *sqlx.DB, productName string) error {
	return nil
}

// UpdateProduct ...
func UpdateProduct(db *sqlx.DB, o *Product) error {
	return nil
}

// DeleteProduct ...
func DeleteProduct(db *sqlx.DB, productID int) error {
	return nil
}

// FindProduct ...
func FindProduct(db *sqlx.DB, productName string) (*Product, error) {
	return nil, nil
}
