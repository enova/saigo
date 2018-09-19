package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func init(){
	Truncate()
}

func createValidProduct(productID int, name string) (*Product, error){
	return NewProduct(GetConn(), productID, name)
}

func TestFindProduct(t *testing.T) {
	defer Truncate()
	productName := "test"
	productNameSubstring := productName[0:2]
	createValidProduct(1, productName)
	product, _  := FindProduct(GetConn(), productNameSubstring)
	assert.NotNil(t, product)
	assert.Equal(t, product.Name, productName)
}
