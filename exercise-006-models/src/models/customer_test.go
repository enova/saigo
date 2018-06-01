package models

import (
"time"
  "database/sql"
  "fmt"
	"testing"
  // "errors"
	// "github.com/stretchr/testify/assert"

  _ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
    var err error
    db, err = sql.Open("postgres", "dbname=test sslmode=disable")
    PanicOn(err)

    _, err = db.Exec("Drop table IF EXISTS customers, products, orders")

    _, err = db.Exec("CREATE TABLE customers(customer_id SERIAL PRIMARY KEY, email text UNIQUE , first_name varchar(70) , last_name varchar(70) , birth_date date , created_at TIMESTAMPTZ default now(), updated_at TIMESTAMPTZ default now())")

    _, err = db.Exec("CREATE TABLE products( product_id SMALLINT PRIMARY KEY, product_name varchar(70))")

    _, err = db.Exec("CREATE TABLE orders(order_id SERIAL PRIMARY KEY, product_id INTEGER NOT NULL REFERENCES products, quantity INTEGER, customer_id INTEGER NOT NULL REFERENCES customers, created_at TIMESTAMPTZ default now(), updated_at TIMESTAMPTZ default now())")
}

func TestSum(t *testing.T) {
    email := "guy@icognito.com"
    first_name := "guy"
    last_name := "icognito"
    bday := time.Now()

    cust, _ := NewCustomer(db, email, first_name, last_name, bday)

    if cust.Email != email {
      message := fmt.Sprintf("Email mismatch: %s != %s ", email, cust.Email)
      t.Fatal(message)
    }

    if cust.FirstName != first_name {
      message := fmt.Sprintf("First_name mismatch: %s != %s ", first_name, cust.FirstName)
      t.Fatal(message)
    }

    if cust.LastName != last_name {
      message := fmt.Sprintf("Last_name mismatch: %s != %s ", last_name, cust.LastName)
      t.Fatal(message)
    }

    if cust.BirthDate != bday {
      message := fmt.Sprintf("Birthday mismatch: %s != %s ", bday, cust.BirthDate)
      t.Fatal(message)
    }
}
