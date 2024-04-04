package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
	"github.com/Hidayathamir/go-product/pkg/goproducterror"
)

//go:generate mockgen -source=product.go -destination=mockusecase/product.go -package=mockusecase

// IProduct contains abstraction of usecase product.
type IProduct interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, req goproductdto.ReqProductSearch) (goproductdto.ResProductSearch, error)
	// GetDetail get product detail by id, or sku, or slug. With priority id > sku > slug.
	GetDetail(ctx context.Context, req goproductdto.ReqProductDetail) (goproductdto.ResProductDetail, error)
	// Add adds product to database.
	Add(ctx context.Context, req goproductdto.ReqProductAdd) (goproductdto.ResProductAdd, error)
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
func (p *Product) Search(ctx context.Context, req goproductdto.ReqProductSearch) (goproductdto.ResProductSearch, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproductdto.ReqProductSearch.Validate: %w", err)
		return goproductdto.ResProductSearch{}, fmt.Errorf("%w: %w", goproducterror.ErrRequestInvalid, err)
	}

	products, err := p.repoProduct.Search(ctx, req.Keyword)
	if err != nil {
		return goproductdto.ResProductSearch{}, fmt.Errorf("Product.repoProduct.Search: %w", err)
	}

	return products, nil
}

// GetDetail implements IProduct.
func (p *Product) GetDetail(ctx context.Context, req goproductdto.ReqProductDetail) (goproductdto.ResProductDetail, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproductdto.ReqProductDetail.Validate: %w", err)
		return goproductdto.ResProductDetail{}, fmt.Errorf("%w: %w", goproducterror.ErrRequestInvalid, err)
	}

	if req.ID != 0 {
		product, err := p.repoProduct.GetDetailByID(ctx, req.ID)
		if err != nil {
			return goproductdto.ResProductDetail{}, fmt.Errorf("Product.repoProduct.GetDetailByID: %w", err)
		}
		return product, nil
	}

	if req.SKU != "" {
		product, err := p.repoProduct.GetDetailBySKU(ctx, req.SKU)
		if err != nil {
			return goproductdto.ResProductDetail{}, fmt.Errorf("Product.repoProduct.GetDetailBySKU: %w", err)
		}
		return product, nil
	}

	if req.Slug != "" {
		product, err := p.repoProduct.GetDetailBySlug(ctx, req.Slug)
		if err != nil {
			return goproductdto.ResProductDetail{}, fmt.Errorf("Product.repoProduct.GetDetailBySlug: %w", err)
		}
		return product, nil
	}

	err = errors.New("request invalid, in theory it's imposible to got this error since it's already handle in req.Validate()")
	return goproductdto.ResProductDetail{}, err
}

// Add implements IProduct.
func (p *Product) Add(ctx context.Context, req goproductdto.ReqProductAdd) (goproductdto.ResProductAdd, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("goproductdto.ReqProductAdd.Validate: %w", err)
		return goproductdto.ResProductAdd{}, fmt.Errorf("%w: %w", goproducterror.ErrRequestInvalid, err)
	}

	productID, err := p.repoProduct.Add(ctx, req)
	if err != nil {
		return goproductdto.ResProductAdd{}, fmt.Errorf("Product.repoProduct.Add: %w", err)
	}

	res := goproductdto.ResProductAdd{
		ID: productID,
	}

	return res, nil
}
