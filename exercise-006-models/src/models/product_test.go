package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init(){
	Truncate()
}

func TestFindProduct(t *testing.T) {
	defer Truncate()
	productName := "test"
	productID := 1

	productNameSubstring := productName[0:2]
	product, err := NewProduct(GetConn(), productID, productName)
	assert.Nil(t, err)
	product, err  = FindProduct(GetConn(), productNameSubstring)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Name, productName)
}
