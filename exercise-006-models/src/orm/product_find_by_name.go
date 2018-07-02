package orm

import "../models"

var qProductFindByName = `
  SELECT
    product_id,
    product_name
  FROM products
  WHERE product_name=$1
  LIMIT 1
`

// FindCustomerByEmail ...
func (cm *ProductORM) FindProductByName(name string) (*models.Product, error) {
  tx := cm.Db.MustBegin()
  row := tx.QueryRow(qProductFindByName, name)

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