package models

import "fmt"

// Order ...
type Product struct {
  ID    int
  Name  string
}

func (p *Product) ToString() string {
  return fmt.Sprintf(
    "(Product#%d %s)",
    p.ID,
    p.Name,
  )
}
