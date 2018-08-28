package models

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func init() {
	conn, err := sqlx.Connect("postgres", "dbname=test sslmode=disable")
	panicOn(err)
	db = conn
}

func TestNewOrder(t *testing.T) {
	// Create a Customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	// Create a product
	product, err := NewProduct(db, 5, "testProduct")
	assert.Nil(t, err)
	defer DeleteProduct(db, product.ID)

	// Create Order
	err = NewOrder(db, customer.ID, product.ID, 2)
	assert.Nil(t, err)

	// Make sure the Order was created in the DB
	retrievedOrder := Order{}
	db.Get(&retrievedOrder, `SELECT * FROM orders WHERE customer_id=$1 AND product_id=$2`, customer.ID, product.ID)
	assert.NotNil(t, retrievedOrder)
	assert.Equal(t, retrievedOrder.CustomerID, customer.ID)
	assert.Equal(t, retrievedOrder.ProductID, product.ID)
	defer DeleteOrder(db, retrievedOrder.ID)
}

func TestUpdateOrder(t *testing.T) {
	// Create a Customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	// Create a product
	product, err := NewProduct(db, 5, "testProduct")
	assert.Nil(t, err)
	defer DeleteProduct(db, product.ID)

	// Create another product
	product2, err := NewProduct(db, 6, "testProduct2")
	assert.Nil(t, err)
	defer DeleteProduct(db, product2.ID)

	// Create Order
	err = NewOrder(db, customer.ID, product.ID, 2)
	assert.Nil(t, err)

	//Retrieve order to get order ID
	var orderID int
	db.Get(&orderID, `SELECT order_id FROM orders WHERE customer_id=$1 AND product_id=$2`, customer.ID, product.ID)

	updatedOrder := Order{ID: orderID, CustomerID: customer.ID, ProductID: product2.ID}
	err = UpdateOrder(db, &updatedOrder)
	assert.Nil(t, err)

	//Retrieve order to verify it was updated
	order := Order{}
	err = db.Get(&order, `SELECT * FROM orders where order_id=$1`, orderID)
	assert.Nil(t, err)
	assert.Equal(t, order.ProductID, product2.ID)

	defer DeleteOrder(db, orderID)
}

func TestDeleteOrder(t *testing.T) {
	// Create a Customer
	customer, err := NewCustomer(db, "test@example.com", "firstname", "lastname", time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
	assert.Nil(t, err)
	defer DeleteCustomer(db, customer.ID)

	// Create a product
	product, err := NewProduct(db, 5, "testProduct")
	assert.Nil(t, err)
	defer DeleteProduct(db, product.ID)

	// Create Order
	err = NewOrder(db, customer.ID, product.ID, 2)
	assert.Nil(t, err)

	//Retrieve order to get order ID
	var orderID int
	db.Get(&orderID, `SELECT order_id FROM orders WHERE customer_id=$1 AND product_id=$2`, customer.ID, product.ID)

	err = DeleteOrder(db, orderID)
	assert.Nil(t, err)

	// There should be no order with ID orderID
	var count int
	db.QueryRowx(`SELECT count(*) FROM orders WHERE order_id=$1`, orderID).Scan(&count)
	assert.Equal(t, count, 0)
}
