package interfaces

import (
	"context"

	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

//go:generate mockgen -source=stock.go -destination=mock/stock.go -package=mock

// UsecaseStock contains abstraction of usecase stock.
type UsecaseStock interface {
	// IncrementStock increment product stock.
	IncrementStock(ctx context.Context, req goproduct.ReqIncrementStock) error
	// DecrementStock decrement product stock.
	DecrementStock(ctx context.Context, req goproduct.ReqDecrementStock) error
}

// RepoStock contains abstraction of repo stock.
type RepoStock interface {
	// IncrementStock increment product stock.
	IncrementStock(ctx context.Context, productID int64) error
	// DecrementStock decrement product stock.
	DecrementStock(ctx context.Context, productID int64) error
}
