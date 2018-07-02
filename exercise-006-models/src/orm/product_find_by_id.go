package orm

import "../models"

var qProductFindById = `
  SELECT
    product_id,
    product_name
  FROM products
  WHERE product_id=$1
  LIMIT 1
`

// FindCustomerByEmail ...
func (cm *ProductORM) FindProductByID(id int) (*models.Product, error) {
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(qProductFindById, id)

  prod := &models.Product{}
  error := row.Scan(
    &prod.ID,
    &prod.Name,
  )
  if error == nil {
    tx.Commit()
    return prod, nil
  } else {
    return nil, error
  }
}