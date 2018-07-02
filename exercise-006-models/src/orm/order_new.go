package orm

import "../models"

var qOrderInsert = `
  INSERT INTO orders (
    customer_id,
    product_id,
    quantity
  ) VALUES (
    $1, $2, $3
  ) RETURNING order_id
`

// NewOrder ...
func (om *OrderORM) NewOrder(
  customerId int,
  productId int,
  quantity int,
) (*models.Order, error) {
  // Build Order object
  ord := &models.Order{
    CustomerID: customerId,
    ProductID:  productId,
    Quantity:   quantity,
  }

  // Attempt to insert
  tx := om.Db.MustBegin()
  row := tx.QueryRow(
    qOrderInsert,
    ord.CustomerID,
    ord.ProductID,
    ord.Quantity,
  )
  error := row.Scan(&ord.ID)

  // Return the result
  if error == nil {
    tx.Commit()
    return ord, nil
  }
  return nil, error
}
