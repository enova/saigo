package orm

var qDelete = `
  DELETE FROM customers
  WHERE customer_id=$1
`

// DeleteCustomer ...
func (cm *CustomerORM) DeleteCustomer(id int) error {
  tx := cm.Db.MustBegin()
  _, error := tx.Exec(qDelete, id)

  // Return the result
  if error == nil {
    tx.Commit()
    return nil
  }
  return error
}