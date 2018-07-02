package orm

import "../models"

var qOrderUpdate = `
  UPDATE orders
  SET 
    customer_id=COALESCE($2, customer_id),
    product_id=COALESCE($3, product_id),
    quantity=COALESCE($4, quantity)
  WHERE order_id=$1
  RETURNING order_id, customer_id, product_id, quantity
`

// Update Order ...
func (om *OrderORM) UpdateOrder(o *models.Order) error {
  // Attempt to update
  tx := om.Db.MustBegin()
  row := tx.QueryRow(
    qOrderUpdate,
    o.ID,
    o.CustomerID,
    o.ProductID,
    o.Quantity,
  )

  error := row.Scan(
    &o.ID,
    &o.CustomerID,
    &o.ProductID,
    &o.Quantity,
  )
  tx.Commit()
  return error
}