package repo

import (
	"context"

	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

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
