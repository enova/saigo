package models

import (
	// "time"

	_ "github.com/lib/pq"
)

// Product ...
type Product struct {
    ID        int
    Product   string
}
//
// func FindProduct(db *sqlx.DB, product string) (*Product, error) {
//   prod, err := db.Query("select * from products where product_name=$1", product)
//   PanicOn(err)
//   return prod, err
// }
