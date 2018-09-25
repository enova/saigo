package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init(){
	Truncate()
}

func createValidOrder(t *testing.T, productID int, quantity int) (*Product, *Customer){
	t.Helper()
	p, _ := NewProduct(GetConn(), productID, "test")
	c, _ := NewCustomer(GetConn(), "test@example.com", "John", "Smith", time.Now())
	err := NewOrder(GetConn(), c.ID, p.ID, quantity)
	assert.Nil(t, err)
	return p, c
}

func TestNewOrder(t *testing.T) {
	defer Truncate()
	p, c := createValidOrder(t, 1,1)
	o, _ := FindMostRecentOrder(GetConn())
	assert.Equal(t, p.ID, o.ProductID)
	assert.Equal(t, c.ID, o.CustomerID)
}

func TestUpdateOrder(t *testing.T) {
	defer Truncate()
	createValidOrder(t, 1,1)
	newQuantity := 100
	o, _ := FindMostRecentOrder(GetConn())
	o.Quantity = newQuantity
	UpdateOrder(GetConn(), &o)
	o, _ = FindMostRecentOrder(GetConn())
	assert.Equal(t, o.Quantity, newQuantity)
}

func TestDeleteOrder(t *testing.T) {
	defer Truncate()
	createValidOrder(t, 1,1)
	o, _ := FindMostRecentOrder(GetConn())
	DeleteOrder(GetConn(), o.ID)
	o, err := FindMostRecentOrder(GetConn())
	assert.NotNil(t, err)
}
