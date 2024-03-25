package usecase

import (
	"context"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

//go:generate mockgen -source=product.go -destination=mockusecase/product.go -package=mockusecase

// IProduct contains abstraction of usecase product.
type IProduct interface {
	// Search search product by name.
	Search(ctx context.Context, req goproduct.ReqProductSearch) (goproduct.ResProductSearch, error)
	// GetDetail get product detail by id, or sku, or slug. With priority id > sku > slug.
	GetDetail(ctx context.Context, req goproduct.ReqProductDetail) (goproduct.ResProductDetail, error)
}

// Product implement IProduct.
type Product struct {
	cfg         config.Config
	repoProduct repo.IProduct
}

var _ IProduct = &Product{}

// NewProduct return *Product which implement IProduct.
func NewProduct(cfg config.Config, repoProduct repo.IProduct) *Product {
	return &Product{
		cfg:         cfg,
		repoProduct: repoProduct,
	}
}

// GetDetail implements IProduct.
func (p *Product) GetDetail(context.Context, goproduct.ReqProductDetail) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}

// Search implements IProduct.
func (p *Product) Search(context.Context, goproduct.ReqProductSearch) (goproduct.ResProductSearch, error) {
	panic("unimplemented")
}
