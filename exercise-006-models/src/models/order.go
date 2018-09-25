package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Order ...
type Order struct {
	ID         int `db:"order_id"`
	ProductID  int `db:"product_id"`
	Quantity   int `db:"quantity"`
	CustomerID int `db:"customer_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// NewOrder ...
func NewOrder(db *sqlx.DB, customerID int, productID int, quantity int) error {
	o := Order{
		ProductID: productID,
		CustomerID: customerID,
		Quantity:quantity,
	}
	query := `INSERT INTO orders (product_id, customer_id, quantity) 
		VALUES(:product_id, :customer_id, :quantity)`
	_, err := db.NamedExec(query, &o)
	return err
}

// Update Order ...
func UpdateOrder(db *sqlx.DB, o *Order) error {
	query := `
		UPDATE orders 
			SET product_id = :product_id, 
				customer_id = :customer_id, 
				quantity = :quantity, 
				updated_at = NOW() 
			WHERE order_id = :order_id`
	_, err := db.NamedExec(query, o)
	return err
}

// DeleteOrder ...
func DeleteOrder(db *sqlx.DB, orderID int) error {
	query := "DELETE FROM orders WHERE order_id = $1"
	_, err := db.Exec(query, orderID)
	return err
}

func FindMostRecentOrder(db *sqlx.DB) (Order, error){
	var o Order
	err := db.Get(&o, "SELECT * FROM orders ORDER BY created_at DESC LIMIT 1")
	return o, err
}
