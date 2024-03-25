package usecase

import (
	"context"

	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

// IProduct contains abstraction of usecase product.
type IProduct interface {
	// Search search product by name.
	Search(ctx context.Context, req goproduct.ReqProductSearch) (goproduct.ResProductSearch, error)
	// GetDetail get product detail by id, or sku, or slug. With priority id > sku > slug.
	GetDetail(ctx context.Context, req goproduct.ReqProductDetail) (goproduct.ResProductDetail, error)
}
