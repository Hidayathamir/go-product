package ctrlgrpc

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
)

// Stock is controller GRPC for stock related.
type Stock struct {
	goproductgrpc.UnimplementedStockServer

	cfg          config.Config
	usecaseStock interfaces.UsecaseStock
}

var _ goproductgrpc.StockServer = &Stock{}

func newStock(cfg config.Config, usecaseStock interfaces.UsecaseStock) *Stock {
	return &Stock{
		cfg:          cfg,
		usecaseStock: usecaseStock,
	}
}

// IncrementStock implements goproductgrpc.StockServer.
func (s *Stock) IncrementStock(c context.Context, r *goproductgrpc.ReqIncrementStock) (*goproductgrpc.StockEmpty, error) {
	req := goproduct.ReqIncrementStock{
		ProductID: r.GetProductId(),
	}

	err := s.usecaseStock.IncrementStock(c, req)
	if err != nil {
		return nil, fmt.Errorf("Stock.usecaseStock.IncrementStock: %w", err)
	}

	res := &goproductgrpc.StockEmpty{}

	return res, nil
}

// DecrementStock implements goproductgrpc.StockServer.
func (s *Stock) DecrementStock(c context.Context, r *goproductgrpc.ReqDecrementStock) (*goproductgrpc.StockEmpty, error) {
	req := goproduct.ReqDecrementStock{
		ProductID: r.GetProductId(),
	}

	err := s.usecaseStock.DecrementStock(c, req)
	if err != nil {
		return nil, fmt.Errorf("Stock.usecaseStock.DecrementStock: %w", err)
	}

	res := &goproductgrpc.StockEmpty{}

	return res, nil
}
