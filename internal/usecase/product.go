package usecase

import (
	"context"
	"errors"
	"fmt"

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

// Search implements IProduct.
func (p *Product) Search(ctx context.Context, req goproduct.ReqProductSearch) (goproduct.ResProductSearch, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproduct.ReqProductSearch.Validate: %w", err)
		return goproduct.ResProductSearch{}, fmt.Errorf("%w: %w", goproduct.ErrRequestInvalid, err)
	}

	products, err := p.repoProduct.SearchByName(ctx, req.Name)
	if err != nil {
		return goproduct.ResProductSearch{}, fmt.Errorf("Product.repoProduct.SearchByName: %w", err)
	}

	return products, nil
}

// GetDetail implements IProduct.
func (p *Product) GetDetail(ctx context.Context, req goproduct.ReqProductDetail) (goproduct.ResProductDetail, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproduct.ReqProductDetail.Validate: %w", err)
		return goproduct.ResProductDetail{}, fmt.Errorf("%w: %w", goproduct.ErrRequestInvalid, err)
	}

	if req.ID != 0 {
		product, err := p.repoProduct.GetDetailByID(ctx, req.ID)
		if err != nil {
			return goproduct.ResProductDetail{}, fmt.Errorf("Product.repoProduct.GetDetailByID: %w", err)
		}
		return product, nil
	}

	if req.SKU != "" {
		product, err := p.repoProduct.GetDetailBySKU(ctx, req.SKU)
		if err != nil {
			return goproduct.ResProductDetail{}, fmt.Errorf("Product.repoProduct.GetDetailBySKU: %w", err)
		}
		return product, nil
	}

	if req.Slug != "" {
		product, err := p.repoProduct.GetDetailBySlug(ctx, req.Slug)
		if err != nil {
			return goproduct.ResProductDetail{}, fmt.Errorf("Product.repoProduct.GetDetailBySlug: %w", err)
		}
		return product, nil
	}

	err = errors.New("request invalid, in theory it's imposible to got this error since it's already handle in req.Validate()")
	return goproduct.ResProductDetail{}, err
}
