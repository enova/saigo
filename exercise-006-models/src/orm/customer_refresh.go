package orm

import "../models"

var qSelect = `
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

/*********************/
// Refresh ...
/*********************/
func (cm *CustomerORM) Refresh(c *models.Customer) error {
  // Attempt to update
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(qSelect, c.ID)
  error := row.Scan(
    &c.ID,
    &c.Email,
    &c.FirstName,
    &c.LastName,
    &c.BirthDate,
  )
  tx.Commit()
  return error
}