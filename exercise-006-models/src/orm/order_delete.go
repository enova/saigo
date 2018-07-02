package orm

var qOrderDelete = `
  DELETE FROM orders
  WHERE order_id=$1
`

// DeleteOrder ...
func (om *OrderORM) DeleteOrder(orderID int) error {
  tx := om.Db.MustBegin()
  _, error := tx.Exec(qOrderDelete, orderID)

  // Return the result
  if error == nil {
    tx.Commit()
    return nil
  }
  return error
}