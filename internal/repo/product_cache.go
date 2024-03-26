package repo

import (
	"context"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

// ProductCache implement IProductCache.
type ProductCache struct {
	cfg config.Config
	// TODO: IMPLEMENT
	// redis redis.Client
}

var _ IProductCache = &ProductCache{}

// NewProductCache return *ProductCache which implement repo.IProductCache.
func NewProductCache(cfg config.Config) *ProductCache {
	return &ProductCache{
		cfg: cfg,
	}
}

// GetDetailByID implements IProductCache.
func (p *ProductCache) GetDetailByID(context.Context, int64) (goproduct.ResProductDetail, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// SetDetailByID implements IProductCache.
func (p *ProductCache) SetDetailByID(context.Context, goproduct.ResProductDetail) error {
	panic("unimplemented") // TODO: IMPLEMENT
}

// GetDetailBySKU implements IProductCache.
func (p *ProductCache) GetDetailBySKU(context.Context, string) (goproduct.ResProductDetail, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// SetDetailBySKU implements IProductCache.
func (p *ProductCache) SetDetailBySKU(context.Context, goproduct.ResProductDetail) error {
	panic("unimplemented") // TODO: IMPLEMENT
}

// GetDetailBySlug implements IProductCache.
func (p *ProductCache) GetDetailBySlug(context.Context, string) (goproduct.ResProductDetail, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// SetDetailBySlug implements IProductCache.
func (p *ProductCache) SetDetailBySlug(context.Context, goproduct.ResProductDetail) error {
	panic("unimplemented") // TODO: IMPLEMENT
}
