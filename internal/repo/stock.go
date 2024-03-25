package repo

import (
	"context"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/db"
)

//go:generate mockgen -source=stock.go -destination=mockrepo/stock.go -package=mockrepo

// IStock contains abstraction of repo stock.
type IStock interface {
	// IncrementStock increment product stock.
	IncrementStock(ctx context.Context, productID int64) error
	// DecrementStock decrement product stock.
	DecrementStock(ctx context.Context, productID int64) error
}

// Stock implement IStock.
type Stock struct {
	cfg config.Config
	db  *db.Postgres
}

var _ IStock = &Stock{}

// NewStock return *Stock which implement repo.IStock.
func NewStock(cfg config.Config, db *db.Postgres) *Stock {
	return &Stock{
		cfg: cfg,
		db:  db,
	}
}

// IncrementStock implements IStock.
func (s *Stock) IncrementStock(context.Context, int64) error {
	panic("unimplemented")
}

// DecrementStock implements IStock.
func (s *Stock) DecrementStock(context.Context, int64) error {
	panic("unimplemented")
}
