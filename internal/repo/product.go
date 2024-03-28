package repo

import (
	"context"

	"github.com/Hidayathamir/go-product/internal/repo/repointerfaces"
	"github.com/Hidayathamir/go-product/internal/usecase/usecaseinterfaces"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
)

type Product struct {
	inMemDB        repointerfaces.ProductInMemoryDB
	noSQLDB        repointerfaces.ProductNoSQLDB
	searchEngineDB repointerfaces.ProductSearchEngineDB
	sqlDB          repointerfaces.ProductSQLDB
}

var _ usecaseinterfaces.RepoProduct = &Product{}

// Search implements usecaseinterfaces.RepoProduct.
func (p *Product) Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error) {
	panic("unimplemented")
}

// GetDetailByID implements usecaseinterfaces.RepoProduct.
func (p *Product) GetDetailByID(ctx context.Context, id int64) (goproduct.ResProductDetail, error) {
	product, err := p.inMemDB.GetDetailByID(ctx, id)
	if err == nil {
		return product, nil
	}

	product, err = p.sqlDB.GetDetailByID(ctx, id)
	if err != nil {
		return goproduct.ResProductDetail{}, err
	}

	err = p.inMemDB.SetDetailByID(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		return goproduct.ResProductDetail{}, err
	}

	return product, nil
}

// GetDetailBySKU implements usecaseinterfaces.RepoProduct.
func (p *Product) GetDetailBySKU(ctx context.Context, sku string) (goproduct.ResProductDetail, error) {
	product, err := p.inMemDB.GetDetailBySKU(ctx, sku)
	if err == nil {
		return product, nil
	}

	product, err = p.sqlDB.GetDetailBySKU(ctx, sku)
	if err != nil {
		return goproduct.ResProductDetail{}, err
	}

	err = p.inMemDB.SetDetailByID(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		return goproduct.ResProductDetail{}, err
	}

	return product, nil
}

// GetDetailBySlug implements usecaseinterfaces.RepoProduct.
func (p *Product) GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error) {
	product, err := p.inMemDB.GetDetailBySlug(ctx, slug)
	if err == nil {
		return product, nil
	}

	product, err = p.sqlDB.GetDetailBySlug(ctx, slug)
	if err != nil {
		return goproduct.ResProductDetail{}, err
	}

	err = p.inMemDB.SetDetailByID(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		return goproduct.ResProductDetail{}, err
	}

	return product, nil
}
