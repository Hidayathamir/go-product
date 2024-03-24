package entity

import "time"

// Stock is entity stock, in db it's table `stock`.
type Stock struct {
	ID        int64
	ProductID int64
	Stock     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
