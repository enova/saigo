package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	conn, err := sqlx.Connect("postgres", "dbname=test sslmode=disable")
	PanicOn(err)
	db = conn
}

func TestNewOrder(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	product, err := NewProduct(db, 5, "bag")
	assert.Nil(t, err)
	defer DeleteProduct(db, product.ID)

	err = NewOrder(db, customer.ID, product.ID, 2)
	assert.Nil(t, err)

	actOrder := Order{}
	db.Get(&actOrder, `SELECT * FROM orders WHERE customer_id=$1 AND product_id=$2`, customer.ID, product.ID)
	assert.NotNil(t, actOrder)
	assert.Equal(t, actOrder.CustomerID, customer.ID)
	assert.Equal(t, actOrder.ProductID, product.ID)
	defer DeleteOrder(db, actOrder.ID)
}

func TestUpdateOrder(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	product, err := NewProduct(db, 5, "bag")
	assert.Nil(t, err)
	defer DeleteProduct(db, product.ID)

	product2, err := NewProduct(db, 6, "gear")
	assert.Nil(t, err)
	defer DeleteProduct(db, product2.ID)

	err = NewOrder(db, customer.ID, product.ID, 2)
	assert.Nil(t, err)

	var orderID int
	db.Get(&orderID, "SELECT order_id FROM orders WHERE customer_id=$1 AND product_id=$2", customer.ID, product.ID)

	updatedOrder := Order{ID: orderID, CustomerID: customer.ID, ProductID: product2.ID}
	err = UpdateOrder(db, &updatedOrder)
	assert.Nil(t, err)

	order := Order{}
	err = db.Get(&order, "SELECT * FROM orders where order_id=$1", orderID)
	assert.Nil(t, err)
	assert.Equal(t, order.ProductID, product2.ID)

	defer DeleteOrder(db, orderID)
}

func TestDeleteOrder(t *testing.T) {
	customer, err := NewCustomer(db, "cust1@test.com", "Bruce", "Leroy", time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	product, err := NewProduct(db, 5, "bag")
	assert.Nil(t, err)
	defer DeleteProduct(db, product.ID)

	err = NewOrder(db, customer.ID, product.ID, 2)
	assert.Nil(t, err)

	//Retrieve order to get order ID
	var orderID int
	db.Get(&orderID, `SELECT order_id FROM orders WHERE customer_id=$1 AND product_id=$2`, customer.ID, product.ID)

	err = DeleteOrder(db, orderID)
	assert.Nil(t, err)

	var count int
	db.QueryRowx(`SELECT count(*) FROM orders WHERE order_id=$1`, orderID).Scan(&count)
	assert.Equal(t, count, 0)
}
