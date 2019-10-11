## Description
A client has asked you to implement a Go package to handle the persistent storage and retrieval of _customers_ and their _orders_.
The client's back-end servers run PostgreSQL 9.x, and a database engineer has already written the required migrations.
Your job is to complete the implementation.


## Engineering Tasks

The client has included a skeleton under the package name `models` that included structures for `Customer`,`Product` and `Order`.

```go
type Customer struct {
    ID        int
    Email     string
    FirstName string
    LastName  string
    BirthDate time.Time
    Orders    []*Order
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Product struct {
    ID        int
    Product   string
}

type Order struct {
    ID         int
    ProductID  int
    Quantity   int
    Time       time.Time
    CustomerID int
}
```

Along with these structures the client has added a collection of stubbed functions:

```go
func (c *Customer) Refresh(db *sql.DB) error

func NewCustomer(db *sql.DB, email, first_name, last_name string, birth_date time.Time) (*Customer, error)
func DeleteCustomer(db *sql.DB, id int) error
func UpdateCustomer(db *sql.DB, u *Customer) error
func FindCustomerByEmail(db *sql.DB, email string) (*Customer, error)
func FindCustomerByID(db *sql.DB, id int) (*Customer, error)
func AllCustomers(db *sql.DB) ([]*Customer, error)

func FindProduct(db *sql.DB, product string) (*Product, error)

func NewOrder(db *sql.DB, custID int, product string, quantity int) error
func UpdateOrder(db *sql.DB, o *Order) error
func DeleteOrder(db *sql.DB, orderID int) error
```

Each of these functions can hit the database hence they all accept a DB-connection as their first argument
(`db *sql.DB`). The function names should be sufficiently descriptive except possibly `Refresh()`. The `Refresh()`
method should update the values of the receiving `Customer` instance with the latest values from the database.
As mentioned early on in the course, try to schedule a session with an instructor to go over the
requirements of this exercise. You don't want to waste time implementing the wrong functionality.

You must:

1. Implement each of these functions
1. Pay attention to errors that might creep up from the database calls.
1. Write tests for the package:
  - `src/models/customer_test.go`
  - `src/models/product_test.go`
  - `src/models/order_test.go`
