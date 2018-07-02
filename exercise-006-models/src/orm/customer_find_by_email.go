package orm

import "../models"

var qFindByEmail = `
  SELECT
    customer_id,
    email,
    first_name,
    last_name,
    birth_date
  FROM customers
  WHERE email=$1
  LIMIT 1
`

// FindCustomerByEmail ...
func (cm *CustomerORM) FindCustomerByEmail(
  email string,
) (*models.Customer, error) {
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(qFindByEmail, email)

  cust := &models.Customer{}
  error := row.Scan(
    &cust.ID,
    &cust.Email,
    &cust.FirstName,
    &cust.LastName,
    &cust.BirthDate,
  )
  if error == nil {
    tx.Commit()
    return cust, nil
  } else {
    return nil, error
  }
}