package main

import (
  "./orm"
  "fmt"
  "time"
  _ "github.com/lib/pq"
)

func main() {
  corm := orm.NewORM()
  cust, err := corm.NewCustomer("a7@b.com", "Richard", "Smith", time.Now())
  if err != nil {
    fmt.Println("There was a problem inserting the user:")
    fmt.Println(err)
  } else {
    fmt.Println("Hello", cust.ToString())
  }

  cust.FirstName = "Peter"
  cust.LastName = "Jenkins"
  err = corm.Refresh(cust)
  if err != nil {
    fmt.Println("There was a problem refreshing the Customer")
    fmt.Println(err)
  } else {
    fmt.Println("Refreshed", cust.ToString())
  }

  cust.FirstName = "Peter"
  cust.LastName = "Jenkins"
  err = corm.UpdateCustomer(cust)
  if err != nil {
    fmt.Println("There was a problem updating the Customer")
    fmt.Println(err)
  } else {
    fmt.Println("Updated", cust.ToString())
  }

  cust2, err := corm.FindCustomerByEmail("a1@b.com")
  if err != nil {
    fmt.Println("A user was not found with that email.")
  } else {
    fmt.Println("Found by email", cust2.ToString())
  }

  cust2, err = corm.FindCustomerByID(5)
  if err != nil {
    fmt.Println("A user was not found with that ID.")
  } else {
    fmt.Println("Found by ID", cust2.ToString())
  }

  err = corm.DeleteCustomer(cust.ID)
  if err != nil {
    fmt.Println("There was a problem deleting the user:")
    fmt.Println(err)
  } else {
    fmt.Println("Deleted", cust.ToString())
  }

  customers, err := corm.AllCustomers()
  if err != nil {
    fmt.Println("Could not retrieve all customers.")
  } else {
    fmt.Println("Here is an index of all customers:")
    for _, customer := range customers {
      fmt.Println("    ", customer.ToString())
    }
  }
}