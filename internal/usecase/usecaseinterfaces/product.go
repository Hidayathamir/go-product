package usecaseinterfaces

import (
	"context"

	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

//go:generate mockgen -source=product.go -destination=mock/product.go -package=mock

// UsecaseProduct contains abstraction of usecase product.
type UsecaseProduct interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, req goproduct.ReqProductSearch) (goproduct.ResProductSearch, error)
	// GetDetail get product detail by id, or sku, or slug. With priority id > sku > slug.
	GetDetail(ctx context.Context, req goproduct.ReqProductDetail) (goproduct.ResProductDetail, error)
}

// RepoProduct contains abstraction of repo product.
type RepoProduct interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error)
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
}
