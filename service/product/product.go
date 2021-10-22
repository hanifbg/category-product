package product

import "time"

type Product struct {
	ID         uint
	CategoryID uint
	Name       string
	Price      int
	Stock      int
	Image      string
	Detail     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
