package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var db *sqlx.DB
var newID int

func init() {
	connection, err := sqlx.Open("postgres", "dbname=test user=postgres password=postgres sslmode=disable")
	PanicOn(err)
	db = connection

	clearTables()

	newC, err := NewCustomer(db, "testOrders@gmail.com", "new", "customer", time.Now())
	PanicOn(err)
	newID = newC.ID
}

func TestNewOrder(t *testing.T) {
	assert := assert.New(t)

	newOrder, err := NewOrder(db, newID, 3, 20)
	assert.Nil(err)
	assert.Equal(3, newOrder.ProductID)
	assert.Equal(20, newOrder.Quantity)
}

func TestUpdateOrder(t *testing.T) {
	assert := assert.New(t)
	var o Order

	newOrder, err := NewOrder(db, newID, 3, 20)
	assert.Nil(err)

	newOrder.ProductID = 4
	newOrder.Quantity = 100
	err = UpdateOrder(db, newOrder)
	assert.Nil(err)

	query := `SELECT * FROM orders WHERE order_id = ($1)`
	err = db.Get(&o, query, newOrder.ID)
	assert.Nil(err)
	assert.Equal(newOrder.ProductID, o.ProductID)
	assert.Equal(newOrder.Quantity, o.Quantity)
}

func TestDeleteOrder(t *testing.T) {
	assert := assert.New(t)
	var count int

	newOrder, err := NewOrder(db, newID, 1, 1000)
	assert.Nil(err)

	query := `SELECT COUNT(*) FROM orders`
	err = db.Get(&count, query)
	assert.Nil(err)
	before := count

	err = DeleteOrder(db, newOrder.ID)
	assert.Nil(err)

	err = db.Get(&count, query)
	assert.Nil(err)
	assert.Equal(before-1, count)
}

func clearTables() {
	db.Exec("DELETE FROM orders")
	db.Exec("DELETE FROM customers")
}

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}
