package repointerfaces

import (
	"context"
	"time"

	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

// ProductInMemoryDB contains abstraction of repo product cache.
type ProductInMemoryDB interface {
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// SetDetailByID set product detail cache by id.
	SetDetailByID(ctx context.Context, data goproduct.ResProductDetail, expire time.Duration) error

	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// SetDetailBySKU get product detail cache by sku.
	SetDetailBySKU(ctx context.Context, data goproduct.ResProductDetail, expire time.Duration) error

	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
	// SetDetailBySlug get product detail cache by slug.
	SetDetailBySlug(ctx context.Context, data goproduct.ResProductDetail, expire time.Duration) error
}

// ProductNoSQLDB -.
type ProductNoSQLDB interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error)
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
}

// ProductSearchEngineDB -.
type ProductSearchEngineDB interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error)
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
}

// ProductSQLDB -.
type ProductSQLDB interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error)
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
}
