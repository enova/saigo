package orm

import (
  "time"
  "../models"
)

var qCustInsert = `
  INSERT INTO customers (
    email,
    first_name,
    last_name,
    birth_date
  ) VALUES (
    $1, $2, $3, $4
  ) RETURNING customer_id
`

/*********************/
// NewCustomer ...
/*********************/
func (cm *CustomerORM) NewCustomer(
  email       string,
  first_name  string,
  last_name   string,
  birth_date  time.Time,
) (*models.Customer, error) {

  // Build Customer object
  cust := &models.Customer{
    Email:      email,
    FirstName:  first_name,
    LastName:   last_name,
    BirthDate:  birth_date,
  }

  // Attempt to insert
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(
    qCustInsert,
    cust.Email,
    cust.FirstName,
    cust.LastName,
    cust.BirthDate,
  )
  error := row.Scan(&cust.ID)

  // Return the result
  if error == nil {
    tx.Commit()
    return cust, nil
  }
  return nil, error
}