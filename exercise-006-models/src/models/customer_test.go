package models

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func init(){
	Truncate()
}

func TestNewCustomer(t *testing.T) {
	defer Truncate()
	email := "test@example.com"
	customer, err := NewCustomer(GetConn(), email, "John", "Smith", time.Now())
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, email)
}

func TestDeleteCustomer(t *testing.T) {
	defer Truncate()
	customer, err := NewCustomer(GetConn(), "test@example.com", "John", "Smith", time.Now())
	assert.Nil(t, err)
	err  = DeleteCustomer(GetConn(), customer.ID)
	assert.Nil(t, err)
	customer, err = FindCustomerByID(GetConn(), customer.ID)
	assert.Nil(t, customer)
	assert.NotNil(t, err)

}

func TestUpdateCustomer(t *testing.T){
	defer Truncate()
	newEmail := "test+new@example.com"
	customer, err := NewCustomer(GetConn(), "testb+existing@example.com", "John", "Smith", time.Now())
	assert.Nil(t, err)
	customer.Email = newEmail

	err = UpdateCustomer(GetConn(), customer)
	assert.Nil(t, err)

	customer, err = FindCustomerByID(GetConn(), customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, newEmail)
}


func TestFindCustomerByEmail(t *testing.T){
	defer Truncate()
	validEmail := "test@example.com"
	_, err := NewCustomer(GetConn(), validEmail, "John", "Smith", time.Now())
	assert.Nil(t, err)

	customer, err := FindCustomerByEmail(GetConn(), validEmail)
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, validEmail)
}

func TestFindCustomerByID(t *testing.T){
	defer Truncate()
	customer, err := NewCustomer(GetConn(), "test@example.com", "John", "Smith", time.Now())
	assert.Nil(t, err)
	newCustomer, err := FindCustomerByID(GetConn(), customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, newCustomer.ID)
}

func TestAllCustomers(t *testing.T){
	defer Truncate()
	customerCount := 3
	for i :=0; i<customerCount ;i++  {
		email := "test+" + strconv.Itoa(i) + "@example.com"
		_, err := NewCustomer(GetConn(), email, "John", "Smith", time.Now())
		assert.Nil(t, err)
	}
	customers, err := AllCustomers(GetConn())
	assert.Nil(t, err)
	assert.Equal(t, len(customers), 3)
}

func TestRefresh(t *testing.T){
	defer Truncate()
	email := "test@example.com"
	customer, err := NewCustomer(GetConn(), email, "John", "Smith", time.Now())
	assert.Nil(t, err)
	customer.Email = ""
	err = customer.Refresh(GetConn())
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, email)
}

