package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
	"github.com/Hidayathamir/go-product/pkg/goproducterror"
)

//go:generate mockgen -source=stock.go -destination=mockusecase/stock.go -package=mockusecase

// IStock contains abstraction of usecase stock.
type IStock interface {
	// IncrementStock increment product stock.
	IncrementStock(ctx context.Context, req goproductdto.ReqIncrementStock) error
	// DecrementStock decrement product stock.
	DecrementStock(ctx context.Context, req goproductdto.ReqDecrementStock) error
}

// Stock implement IStock.
type Stock struct {
	cfg       config.Config
	repoStock repo.IStock
}

var _ IStock = &Stock{}

// NewStock return *Stock which implement IStock.
func NewStock(cfg config.Config, repoStock repo.IStock) *Stock {
	return &Stock{
		cfg:       cfg,
		repoStock: repoStock,
	}
}

// IncrementStock implements IStock.
func (s *Stock) IncrementStock(ctx context.Context, req goproductdto.ReqIncrementStock) error {
	err := req.Validate()
	if err != nil {
		err := trace.Wrap(err)
		return fmt.Errorf("%w: %w", goproducterror.ErrRequestInvalid, err)
	}

	err = s.repoStock.IncrementStock(ctx, req.ProductID)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

// DecrementStock implements IStock.
func (s *Stock) DecrementStock(ctx context.Context, req goproductdto.ReqDecrementStock) error {
	err := req.Validate()
	if err != nil {
		err := trace.Wrap(err)
		return fmt.Errorf("%w: %w", goproducterror.ErrRequestInvalid, err)
	}

	err = s.repoStock.DecrementStock(ctx, req.ProductID)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}
