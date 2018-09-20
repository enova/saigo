package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db *sqlx.DB

func init() {
	conn, err := sqlx.Connect("postgres", "dbname=test sslmode=disable")
	PanicOn(err)
	db = conn
}

func TestNewProduct(t *testing.T) {
	expProdId := 5
	expProdName := "bag"

	product, err := NewProduct(db, expProdId, expProdName)
	assert.Nil(t, err)
	defer DeleteProduct(db, expProdId)

	assert.Equal(t, product.Name, expProdName)

	var actProd Product
	db.Get(&actProd, `SELECT * FROM products WHERE product_name=$1`, expProdName)
	assert.Equal(t, actProd.Name, expProdName)
}

func TestFindProduct(t *testing.T) {
	expProdId := 5
	expProdName := "bag"

	_, err := NewProduct(db, expProdId, expProdName)
	assert.Nil(t, err)
	defer DeleteProduct(db, expProdId)

	actProd, err := FindProduct(db, expProdName)
	assert.Nil(t, err)
	assert.Equal(t, actProd.Name, expProdName)
}
