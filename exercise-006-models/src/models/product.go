package models

import "github.com/jmoiron/sqlx"

// Product ...
type Product struct {
	ID int `db:"product_id"`
	Name string `db:"product_name"`
}

func FindProduct(db *sqlx.DB, product string) (*Product, error){
	var p Product
	query := "SELECT * FROM products WHERE product_name ILIKE $1"
	err := db.Get(&p, query, "%"+product+"%")
	if err != nil {
		return nil, err
	}
	return &p, err
}


func NewProduct(db *sqlx.DB, productID int, product string) (*Product, error){
	var p Product
	query := "INSERT INTO products (product_id, product_name) VALUES($1, $2) RETURNING *"
	err := db.Get(&p, query, productID, product)
	if err != nil {
		return nil, err
	}
	return &p, err
}

