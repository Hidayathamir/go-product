package entity

import "time"

// Product is entity product, in db it's table `product`.
type Product struct {
	ID        int64
	SKU       string
	Slug      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
