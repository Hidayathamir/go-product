package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

// Stock implement IStock.
type Stock struct {
	cfg       config.Config
	repoStock interfaces.RepoStock
}

var _ interfaces.UsecaseStock = &Stock{}

// NewStock return *Stock which implement IStock.
func NewStock(cfg config.Config, repoStock interfaces.RepoStock) *Stock {
	return &Stock{
		cfg:       cfg,
		repoStock: repoStock,
	}
}

// IncrementStock implements IStock.
func (s *Stock) IncrementStock(ctx context.Context, req goproduct.ReqIncrementStock) error {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproduct.ReqIncrementStock.Validate: %w", err)
		return fmt.Errorf("%w: %w", goproduct.ErrRequestInvalid, err)
	}

	err = s.repoStock.IncrementStock(ctx, req.ProductID)
	if err != nil {
		return fmt.Errorf("Stock.repoStock.IncrementStock: %w", err)
	}

	return nil
}

// DecrementStock implements IStock.
func (s *Stock) DecrementStock(ctx context.Context, req goproduct.ReqDecrementStock) error {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproduct.ReqDecrementStock.Validate: %w", err)
		return fmt.Errorf("%w: %w", goproduct.ErrRequestInvalid, err)
	}

	err = s.repoStock.DecrementStock(ctx, req.ProductID)
	if err != nil {
		return fmt.Errorf("Stock.repoStock.DecrementStock: %w", err)
	}

	return nil
}
