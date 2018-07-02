package main

import (
  "./orm"
  "fmt"
  "time"
  _ "github.com/lib/pq"
)

func exhibitCustomer() {
  corm := orm.ForCustomer()
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

func exhibitOrders() {
  corm := orm.ForCustomer()
  cust, err := corm.FindCustomerByID(1)

  oorm := orm.ForOrder()
  ord, err := oorm.NewOrder(cust.ID, 1, 15)
  if err != nil {
    fmt.Println("There was a problem inserting the order:")
    fmt.Println(err)
  } else {
    fmt.Println("New Order Placed:", ord.ToString())
  }

  ord.Quantity = 78
  err = oorm.UpdateOrder(ord)
  if err != nil {
    fmt.Println("There was a problem updating the order:")
    fmt.Println(err)
  } else {
    fmt.Println("Order Updated:", ord.ToString())
  }

  err = oorm.DeleteOrder(ord.ID)
  if err != nil {
    fmt.Println("There was a problem deleting the order:")
    fmt.Println(err)
  } else {
    fmt.Println("Order Deleted:", ord.ToString())
  }
}

func exhibitProducts() {
  porm := orm.ForProduct()
  prod, err := porm.FindProductByName("kayak")
  if err != nil {
    fmt.Println("There was a problem finding the product:")
    fmt.Println(err)
  } else {
    fmt.Println("Product Found:", prod.ToString())
  }

  prod, err = porm.FindProductByID(4)
  if err != nil {
    fmt.Println("There was a problem finding the product:")
    fmt.Println(err)
  } else {
    fmt.Println("Product Found:", prod.ToString())
  }
}

func main() {
  exhibitCustomer()
  exhibitOrders()
  exhibitProducts()
}