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

func createValidCustomer(email string) (*Customer, error){
	return NewCustomer(GetConn(), email, "John", "Smith", time.Now())
}


func TestNewCustomer(t *testing.T) {
	defer Truncate()
	email := "test@example.com"
	customer, _ := createValidCustomer(email)
	assert.Equal(t, customer.Email, email)
}

func TestDeleteCustomer(t *testing.T) {
	defer Truncate()
	invalidEmail := "test+invalid@example.com"
	customer, _ := createValidCustomer("test@example.com")
	customer.Email = invalidEmail

	DeleteCustomer(GetConn(), customer.ID)
	err := customer.Refresh(GetConn())
	assert.NotNil(t, err)

	// Test that refresh fails, thus customer email remains unchanged
	assert.Equal(t, customer.Email, invalidEmail)
}

func TestUpdateCustomer(t *testing.T){
	defer Truncate()
	invalidEmail := "testa+invalid@example.com"
	validNewEmail := "test+new@example.com"
	customer, _ := createValidCustomer("testb+existing@example.com")
	customer.Email = validNewEmail

	UpdateCustomer(GetConn(), customer)
	customer.Email = invalidEmail
	err := customer.Refresh(GetConn())
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, validNewEmail)
}


func TestFindCustomerByEmail(t *testing.T){
	defer Truncate()
	validEmail := "test@example.com"
	createValidCustomer(validEmail)
	customer, err := FindCustomerByEmail(GetConn(), validEmail)
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, validEmail)
}

func TestFindCustomerByID(t *testing.T){
	defer Truncate()
	customer, _ := createValidCustomer("test@example.com")
	newCustomer, err := FindCustomerByID(GetConn(), customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, newCustomer.ID)
}

func TestAllCustomers(t *testing.T){
	defer Truncate()
	customerCount := 3
	for i :=0; i<customerCount ;i++  {
		createValidCustomer("test+" + strconv.Itoa(i) + "@example.com")
	}
	customers, err := AllCustomers(GetConn())
	assert.Nil(t, err)
	assert.Equal(t, len(customers), 3)
}
