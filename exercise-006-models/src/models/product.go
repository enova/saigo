package models

import (
	"github.com/jmoiron/sqlx"
)

// Product ...
type Product struct {
	ID   int    `db:"product_id"`
	Name string `db:"product_name"`
}

// NewProduct ...
func NewProduct(db *sqlx.DB, id int, name string) (*Product, error) {
	var p Product
	query := `INSERT INTO products(product_id, product_name) VALUES ($1, $2) RETURNING *`
	err := db.Get(&p, query, id, name)
	return &p, err
}

// DeleteProduct ...
func DeleteProduct(db *sqlx.DB, productID int) error {
	query := `DELETE FROM products WHERE product_id = ($1)`
	_, err := db.Exec(query, productID)
	return err
}
