package models

import (
	"time"
  "database/sql"

	_ "github.com/lib/pq"
)

// Customer ...
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

// // Refresh ...
// func Refresh(db *sql.DB, c *Customer) error {
//   cust, err := FindCustomerByID(db, c.ID)
//   PanicOn(err)
//   c = cust
// 	return err
// }

// NewCustomer ...
func NewCustomer(db *sql.DB, email string, first_name string, last_name string, birth_date time.Time) (*Customer, error) {
  result, err := db.Exec("Insert into customers(email, first_name, last_name, birth_date) values($1, $2, $3, $4)", email, first_name, last_name, birth_date)
  PanicOn(err)

  id, err := result.LastInsertId()

  cust, err := FindCustomerByID(db, int(id))



  return cust, err
}

// // DeleteCustomer ...
// func DeleteCustomer(db *sql.DB, id int) error {
//   err := db.Exec("Delete from customers where customer_id = $1", id)
//   PanicOn(err)
// 	return err
// }
//
// // UpdateCustomer ...
// func UpdateCustomer(db *sql.DB, u *Customer) error {
//   err := db.Exec("update customers set email=$1, first_name=$2, last_name=$3, birth_date=$4 where customer_id=$5", u.email, u.first_name, u.last_name, u.birth_date, u.customer_id)
//   PanicOn(err)
//   return err
// }
//
// // FindCustomerByEmail ...
// func FindCustomerByEmail(db *sql.DB, email string) (*Customer, error) {
//   cust, err := db.Query("Select * from customers where email=$1", email)
//   PanicOn(err)
// 	return cust, err
// }

// FindCustomerByID ...
func FindCustomerByID(db *sql.DB, id int) (*Customer, error) {
  result, err := db.Query("Select * from customers where customer_id=$1", id)
  PanicOn(err)

  var cust *Customer
  for result.Next() {
    var id int
    var email string
    var fname string
    var lname string
    var bdate time.Time
    var created time.Time
    var updated time.Time

    err := result.Scan(&id, &email, &fname, &lname, &bdate, &created, &updated)
    PanicOn(err)

    cust = &Customer{id, email, fname, lname, bdate, nil, created, updated}
  }
  return cust, err
}

// // AllCustomers ...
// func AllCustomers(db *sql.DB) ([]*Customer, error) {
//   cust, err := db.Query("Select * from customers")
//   PanicOn(err)
//   return cust, err
// }
