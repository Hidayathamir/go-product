package ctrlgrpc

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Product is controller GRPC for product related.
type Product struct {
	goproductgrpc.UnimplementedProductServer

	cfg            config.Config
	usecaseProduct interfaces.UsecaseProduct
}

var _ goproductgrpc.ProductServer = &Product{}

func newProduct(cfg config.Config, usecaseProduct interfaces.UsecaseProduct) *Product {
	return &Product{
		cfg:            cfg,
		usecaseProduct: usecaseProduct,
	}
}

// Search implements goproductgrpc.ProductServer.
func (p *Product) Search(c context.Context, r *goproductgrpc.ReqProductSearch) (*goproductgrpc.ResProductSearch, error) {
	req := goproduct.ReqProductSearch{
		Keyword: r.GetKeyword(),
	}

	resSearch, err := p.usecaseProduct.Search(c, req)
	if err != nil {
		err := fmt.Errorf("Product.usecaseProduct.Search: %w", err)
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
	req := goproduct.ReqProductDetail{
		ID:   r.GetId(),
		SKU:  r.GetSku(),
		Slug: r.GetSlug(),
	}

	resGetDetail, err := p.usecaseProduct.GetDetail(c, req)
	if err != nil {
		err := fmt.Errorf("Product.usecaseProduct.GetDetail: %w", err)
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
