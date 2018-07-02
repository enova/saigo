package orm

import "../models"

var qFindById = `
  SELECT
    customer_id,
    email,
    first_name,
    last_name,
    birth_date
  FROM customers
  WHERE customer_id=$1
  LIMIT 1
`

// FindCustomerByID ...
func (cm *CustomerORM) FindCustomerByID(id int) (*models.Customer, error) {
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(qFindById, id)

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