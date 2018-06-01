package models

import (
	"time"

	_ "github.com/lib/pq"
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

// // NewOrder ...
// func NewOrder(db *sql.DB, customerID int, productID int, quantity int) error {
//   err := db.Exec("Insert into orders (product_id, quantity, customer_id) values ($1, $2, $3)", product_id, quantity, customerID)
//   PanicOn(err)
//   return err
// }
//
// // Update Order ...
// func UpdateOrder(db *sql.DB, o *Order) error {
//   err := db.Exec("Update oreders set product_id=$1, quantity=$2, customer_id=$3 where order_id=$4", o.productID, o.quantity, o.customerID, o.ID)
//   PanicOn(err)
//   return err
// }
//
// // DeleteOrder ...
// func DeleteOrder(db *sql.DB, orderID int) error {
//   err := db.Exec("Delete from orders where order_id=$1", orderID)
//   PanicOn(err)
// 	return err
// }
