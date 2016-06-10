## Description
A client has asked you to implement a Go package to handle the persistent storage and retrieval of _customers_ and there _orders_.
The client's back-end servers run PostgreSQL 9.x, and a database engineer has already written the required migrations.
Your job is to complete the implementation.


## Engineering Tasks

The client has included a skeleton under the package name `models` that included structures for `Customer`,`Product` and `Order`.

```
type Customer struct {
    ID        int
    Email     string
    FirstName string
    LastName  string
    BirthDate date
    Orders    []*Order
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

```
func (c *Customer) Refresh(db *sqlx.DB) error

func NewCustomer(db *sqlx.DB, email, first, last string, born time.Time) (*Customer, error)
func DeleteCustomer(db *sqlx.DB, id int) error
func UpdateCustomer(db *sqlx.DB, u *Customer) error
func FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error)
func FindCustomerByID(db *sqlx.DB, id int) (*Customer, error)
func AllCustomers(db *sqlx.DB) ([]*Customer, error)

func FindProduct(db *sqlx.DB, product string) (*Product, error)

func NewOrder(db *sqlx.DB, custID int, product string, quantity int) error
func UpdateOrder(db *sqlx.DB, o *Order) error
func DeleteOrder(db *sqlx.DB, orderID int) error
```

Each of these functions can hit the database hence they all accept a DB-connection as their first argument
(`db *sqlx.DB`). The function names should be sufficiently descriptive except possibly `Refresh()`. The `Refresh()`
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

## Do I have to use `sqlx`?
No. If you prefer not to use the [sqlx](http://github.com/jmoiron/sqlx) package and would instead prefer to work with
the standard `database/sql` package, that is completely fine. Please feel free to replace all instances of `sqlx` in the
code with `sql` and code away!
