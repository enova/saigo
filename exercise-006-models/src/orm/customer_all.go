package orm

import "../models"

var qSelectAll = `
  SELECT
    customer_id,
    email,
    first_name,
    last_name,
    birth_date
  FROM customers
`

// AllCustomers ...
func (cm *CustomerORM) AllCustomers() ([]*models.Customer, error) {
  customers := make([]*models.Customer, 0)

  tx := cm.Db.MustBegin()
  rows, err := tx.Query(qSelectAll)
  if err != nil {
    return customers, err
  }
  defer rows.Close()

  for rows.Next() {
    cust := &models.Customer{}
    error := rows.Scan(
      &cust.ID,
      &cust.Email,
      &cust.FirstName,
      &cust.LastName,
      &cust.BirthDate,
    )
    if error == nil {
      customers = append(customers, cust)
    }
  }

  return customers, nil
}