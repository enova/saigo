package models

import (
	"testing"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var db *sqlx.DB

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	conn, err := sqlx.Connect("postgres", "dbname=test sslmode=disable")
	panicOn(err)
	db = conn
}

func TestCreateProduct(t *testing.T) {
	testProductID := 5
	testProductName := "testProductName"
	product, err := NewProduct(db, testProductID, testProductName)
	assert.Nil(t, err)
	defer DeleteProduct(db, testProductID)

	assert.Equal(t, product.Name, testProductName)

	// make sure the product got created in DB
	var retrievedProduct Product
	db.Get(&retrievedProduct, `SELECT * FROM products WHERE product_name=$1`, testProductName)
	assert.Equal(t, retrievedProduct.Name, testProductName)
}

func TestFindProduct(t *testing.T) {
	testProductID := 5
	testProductName := "testProductName"

	// Create a product
	_, err := NewProduct(db, testProductID, testProductName)
	assert.Nil(t, err)
	defer DeleteProduct(db, testProductID)

	// Find product by name
	retrievedProdct, err := FindProduct(db, testProductName)
	assert.Nil(t, err)
	assert.Equal(t, retrievedProdct.Name, testProductName)
}
