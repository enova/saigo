package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	connection, err := sqlx.Open("postgres", "dbname=test user=postgres password=postgres sslmode=disable")
	PanicOn(err)
	db = connection
}

func TestNewProduct(t *testing.T) {
	assert := assert.New(t)
	var tempID int

	query := `SELECT product_id FROM products ORDER BY product_id DESC LIMIT 1`
	err := db.Get(&tempID, query)

	newProduct, err := NewProduct(db, tempID+100, "new product")
	assert.Nil(err)
	assert.Equal("new product", newProduct.Name)
}

func TestDeleteProduct(t *testing.T) {
	assert := assert.New(t)
	var count int
	var tempID int

	query := `SELECT product_id FROM products ORDER BY product_id DESC LIMIT 1`
	err := db.Get(&tempID, query)

	newProduct, err := NewProduct(db, tempID+100, "delete product")
	assert.Nil(err)

	query = `SELECT COUNT(*) FROM products`
	err = db.Get(&count, query)
	assert.Nil(err)
	before := count

	err = DeleteProduct(db, newProduct.ID)
	assert.Nil(err)

	err = db.Get(&count, query)
	assert.Nil(err)
	assert.Equal(before-1, count)
}
