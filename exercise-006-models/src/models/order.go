package models

import (
	"time"
	"fmt"
)

// Order ...
type Order struct {
	ID         int
	ProductID  int
	Quantity   int
	CustomerID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o *Order) ToString() string {
	return fmt.Sprintf(
		"(Order#%d for Customer#%d: Product#%d [x%d])",
		o.ID,
		o.CustomerID,
		o.ProductID,
		o.Quantity,
	)
}
