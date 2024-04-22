package controllergrpc

import (
	"context"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
)

// Stock is controller GRPC for stock related.
type Stock struct {
	goproductgrpc.UnimplementedStockServer

	cfg          config.Config
	usecaseStock usecase.IStock
}

var _ goproductgrpc.StockServer = &Stock{}

// NewStock -.
func NewStock(cfg config.Config, usecaseStock usecase.IStock) *Stock {
	return &Stock{
		cfg:          cfg,
		usecaseStock: usecaseStock,
	}
}

// IncrementStock implements goproductgrpc.StockServer.
func (s *Stock) IncrementStock(c context.Context, r *goproductgrpc.ReqIncrementStock) (*goproductgrpc.StockVoid, error) {
	req := goproductdto.ReqIncrementStock{
		ProductID: r.GetProductId(),
	}

	err := s.usecaseStock.IncrementStock(c, req)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	res := &goproductgrpc.StockVoid{}

	return res, nil
}

// DecrementStock implements goproductgrpc.StockServer.
func (s *Stock) DecrementStock(c context.Context, r *goproductgrpc.ReqDecrementStock) (*goproductgrpc.StockVoid, error) {
	req := goproductdto.ReqDecrementStock{
		ProductID: r.GetProductId(),
	}

	err := s.usecaseStock.DecrementStock(c, req)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	res := &goproductgrpc.StockVoid{}

	return res, nil
}
