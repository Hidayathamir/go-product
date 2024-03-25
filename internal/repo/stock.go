package repo

import "context"

// IStock contains abstraction of repo stock.
type IStock interface {
	// IncrementStock increment product stock.
	IncrementStock(ctx context.Context, productID int64) error
	// DecrementStock decrement product stock.
	DecrementStock(ctx context.Context, productID int64) error
}
