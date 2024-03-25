package usecase

import (
	"context"

	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

// IStock contains abstraction of usecase stock.
type IStock interface {
	// IncrementStock increment product stock.
	IncrementStock(ctx context.Context, req goproduct.ReqIncrementStock) error
	// DecrementStock decrement product stock.
	DecrementStock(ctx context.Context, req goproduct.ReqDecrementStock) error
}
