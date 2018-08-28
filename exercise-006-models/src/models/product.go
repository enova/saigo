package models

import "github.com/jmoiron/sqlx"

// Product ...
type Product struct {
	ID   int    `db:"product_id"`
	Name string `db:"product_name"`
}

// FindProduct ...
func FindProduct(db *sqlx.DB, product string) (*Product, error) {
	p := &Product{}
	err := db.Get(p, `SELECT * FROM products WHERE product_name=$1`, product)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// NewProduct ...
func NewProduct(db *sqlx.DB, id int, name string) (*Product, error) {
	var productID int
	err := db.QueryRowx(`INSERT INTO products(product_id, product_name) VALUES($1, $2) RETURNING product_id`, id, name).Scan(&productID)
	if err != nil {
		return nil, err
	}

	product := &Product{}
	err = db.Get(product, `SELECT * FROM products WHERE product_id=$1`, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct ...
func DeleteProduct(db *sqlx.DB, id int) error {
	_, err := db.Exec(`DELETE FROM products WHERE product_id=$1`, id)
	return err
}
