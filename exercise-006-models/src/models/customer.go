package models

import (
	"time"
	"fmt"
)

// Customer ...
type Customer struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	BirthDate time.Time
	Orders    []*Order
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Customer) FullName() string {
	return c.FirstName + " " + c.LastName
}

func (c *Customer) ToString() string {
	return fmt.Sprintf(
		"(%d: %s) Email: %s, Born: %s",
		c.ID,
		c.FullName(),
		c.Email,
		c.BirthDate.Format("2006-01-02"),
	)
}

