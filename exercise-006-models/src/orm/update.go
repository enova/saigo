package orm

import "../models"

var qUpdate = `
  UPDATE customers
  SET 
    email=COALESCE($2, email),
    first_name=COALESCE($3, first_name),
    last_name=COALESCE($4, last_name),
    birth_date=COALESCE($5, birth_date)
  WHERE customer_id=$1
  RETURNING customer_id, email, first_name, last_name, birth_date
`

// UpdateCustomer ...
func (cm *CustomerORM) UpdateCustomer(c *models.Customer) error {
  // Attempt to update
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(
    qUpdate,
    c.ID,
    c.Email,
    c.FirstName,
    c.LastName,
    c.BirthDate,
  )

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