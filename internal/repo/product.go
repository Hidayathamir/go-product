package repo

import (
	"context"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

//go:generate mockgen -source=product.go -destination=mockrepo/product.go -package=mockrepo

// IProduct contains abstraction of repo product.
type IProduct interface {
	// SearchByName search product by name.
	SearchByName(ctx context.Context, name string) (goproduct.ResProductSearch, error)
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
}

// Product implement IProduct.
type Product struct {
	cfg config.Config
	db  *db.Postgres
}

var _ IProduct = &Product{}

// NewProduct return *Product which implement repo.IProduct.
func NewProduct(cfg config.Config, db *db.Postgres) *Product {
	return &Product{
		cfg: cfg,
		db:  db,
	}
}

// SearchByName implements IProduct.
func (p *Product) SearchByName(context.Context, string) (goproduct.ResProductSearch, error) {
	panic("unimplemented")
}

// GetDetailByID implements IProduct.
func (p *Product) GetDetailByID(context.Context, int64) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}

// GetDetailBySKU implements IProduct.
func (p *Product) GetDetailBySKU(context.Context, string) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}

// GetDetailBySlug implements IProduct.
func (p *Product) GetDetailBySlug(context.Context, string) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}
