package grpc

import (
	"context"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Product is transport GRPC for product related.
type Product struct {
	goproductgrpc.UnimplementedProductServer

	cfg            config.Config
	usecaseProduct usecase.IProduct
}

var _ goproductgrpc.ProductServer = &Product{}

// NewProduct -.
func NewProduct(cfg config.Config, usecaseProduct usecase.IProduct) *Product {
	return &Product{
		cfg:            cfg,
		usecaseProduct: usecaseProduct,
	}
}

// Search implements goproductgrpc.ProductServer.
func (p *Product) Search(c context.Context, r *goproductgrpc.ReqProductSearch) (*goproductgrpc.ResProductSearch, error) {
	req := goproductdto.ReqProductSearch{
		Keyword: r.GetKeyword(),
	}

	resSearch, err := p.usecaseProduct.Search(c, req)
	if err != nil {
		err := trace.Wrap(err)
		return nil, err
	}

	resProducts := []*goproductgrpc.ResProductDetail{}

	for _, product := range resSearch.Products {
		resProducts = append(resProducts, &goproductgrpc.ResProductDetail{
			Id:          product.ID,
			Sku:         product.SKU,
			Slug:        product.Slug,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		})
	}

	res := &goproductgrpc.ResProductSearch{
		Products: resProducts,
	}

	return res, nil
}

// GetDetail implements goproductgrpc.ProductServer.
func (p *Product) GetDetail(c context.Context, r *goproductgrpc.ReqProductDetail) (*goproductgrpc.ResProductDetail, error) {
	req := goproductdto.ReqProductDetail{
		ID:   r.GetId(),
		SKU:  r.GetSku(),
		Slug: r.GetSlug(),
	}

	resGetDetail, err := p.usecaseProduct.GetDetail(c, req)
	if err != nil {
		err := trace.Wrap(err)
		return nil, err
	}

	res := &goproductgrpc.ResProductDetail{
		Id:          resGetDetail.ID,
		Sku:         resGetDetail.SKU,
		Slug:        resGetDetail.Slug,
		Name:        resGetDetail.Name,
		Description: resGetDetail.Description,
		Stock:       resGetDetail.Stock,
		CreatedAt:   timestamppb.New(resGetDetail.CreatedAt),
		UpdatedAt:   timestamppb.New(resGetDetail.UpdatedAt),
	}

	return res, nil
}

// Add implements goproductgrpc.ProductServer.
func (p *Product) Add(c context.Context, r *goproductgrpc.ReqProductAdd) (*goproductgrpc.ResProductAdd, error) {
	req := goproductdto.ReqProductAdd{
		SKU:         r.GetSku(),
		Slug:        r.GetSlug(),
		Name:        r.GetName(),
		Description: r.GetDescription(),
		Stock:       r.GetStock(),
	}

	resAdd, err := p.usecaseProduct.Add(c, req)
	if err != nil {
		err := trace.Wrap(err)
		return nil, err
	}

	res := &goproductgrpc.ResProductAdd{
		Id: resAdd.ID,
	}

	return res, nil
}
