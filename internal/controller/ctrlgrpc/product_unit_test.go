package ctrlgrpc

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces/mock"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUnitProductSearch(t *testing.T) {
	t.Parallel()

	t.Run("product search success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProduct := mock.NewMockUsecaseProduct(ctrl)

		p := &Product{
			cfg:            config.Config{},
			usecaseProduct: usecaseProduct,
		}

		reqUsecase := goproduct.ReqProductSearch{
			Keyword: "iphone13",
		}

		timeNow := time.Now()
		resUsecaseProduct1 := goproduct.ResProductDetail{
			ID:          2342,
			SKU:         "asdfasd",
			Slug:        "afesfes",
			Name:        "sefes",
			Description: "sefesa",
			Stock:       2342,
			CreatedAt:   timeNow,
			UpdatedAt:   timeNow,
		}
		resUsecaseProduct2 := goproduct.ResProductDetail{
			ID:          232,
			SKU:         "fasefv",
			Slug:        "esreq",
			Name:        "eefqadf",
			Description: "weqerqwer",
			Stock:       234,
			CreatedAt:   timeNow,
			UpdatedAt:   timeNow,
		}
		resUsecaseProductSearch := goproduct.ResProductSearch{
			Products: []goproduct.ResProductDetail{resUsecaseProduct1, resUsecaseProduct2},
		}

		usecaseProduct.EXPECT().
			Search(context.Background(), reqUsecase).
			Return(resUsecaseProductSearch, nil)

		req := &goproductgrpc.ReqProductSearch{Keyword: reqUsecase.Keyword}

		res, err := p.Search(context.Background(), req)

		resGRPCProduct1 := &goproductgrpc.ResProductDetail{
			Id:          resUsecaseProduct1.ID,
			Sku:         resUsecaseProduct1.SKU,
			Slug:        resUsecaseProduct1.Slug,
			Name:        resUsecaseProduct1.Name,
			Description: resUsecaseProduct1.Description,
			Stock:       resUsecaseProduct1.Stock,
			CreatedAt:   timestamppb.New(resUsecaseProduct1.CreatedAt),
			UpdatedAt:   timestamppb.New(resUsecaseProduct1.UpdatedAt),
		}
		resGRPCProduct2 := &goproductgrpc.ResProductDetail{
			Id:          resUsecaseProduct2.ID,
			Sku:         resUsecaseProduct2.SKU,
			Slug:        resUsecaseProduct2.Slug,
			Name:        resUsecaseProduct2.Name,
			Description: resUsecaseProduct2.Description,
			Stock:       resUsecaseProduct2.Stock,
			CreatedAt:   timestamppb.New(resUsecaseProduct2.CreatedAt),
			UpdatedAt:   timestamppb.New(resUsecaseProduct2.UpdatedAt),
		}
		resGRPCProductSearch := &goproductgrpc.ResProductSearch{
			Products: []*goproductgrpc.ResProductDetail{resGRPCProduct1, resGRPCProduct2},
		}

		require.NoError(t, err)
		assert.Equal(t, resGRPCProductSearch, res)
	})
	t.Run("product search usecaes error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProduct := mock.NewMockUsecaseProduct(ctrl)

		p := &Product{
			cfg:            config.Config{},
			usecaseProduct: usecaseProduct,
		}

		reqUsecase := goproduct.ReqProductSearch{
			Keyword: "iphone13",
		}

		resUsecaseProductSearch := goproduct.ResProductSearch{}

		usecaseProduct.EXPECT().
			Search(context.Background(), reqUsecase).
			Return(resUsecaseProductSearch, assert.AnError)

		req := &goproductgrpc.ReqProductSearch{Keyword: reqUsecase.Keyword}

		res, err := p.Search(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}

func TestUnitProductGetDetail(t *testing.T) {
	t.Parallel()

	t.Run("get product success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProduct := mock.NewMockUsecaseProduct(ctrl)

		p := &Product{
			cfg:            config.Config{},
			usecaseProduct: usecaseProduct,
		}

		reqUsecase := goproduct.ReqProductDetail{ID: 23423}
		resUsecase := goproduct.ResProductDetail{
			ID:          23423,
			SKU:         "asdfaesf",
			Slug:        "fsefsae",
			Name:        "sefaef",
			Description: "aefax",
			Stock:       23423,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		usecaseProduct.EXPECT().
			GetDetail(context.Background(), reqUsecase).
			Return(resUsecase, nil)

		reqGRPC := &goproductgrpc.ReqProductDetail{Id: reqUsecase.ID}
		expectedResGRPC := &goproductgrpc.ResProductDetail{
			Id:          resUsecase.ID,
			Sku:         resUsecase.SKU,
			Slug:        resUsecase.Slug,
			Name:        resUsecase.Name,
			Description: resUsecase.Description,
			Stock:       resUsecase.Stock,
			CreatedAt:   timestamppb.New(resUsecase.CreatedAt),
			UpdatedAt:   timestamppb.New(resUsecase.UpdatedAt),
		}

		res, err := p.GetDetail(context.Background(), reqGRPC)

		require.NoError(t, err)
		assert.Equal(t, expectedResGRPC, res)
	})
	t.Run("get product usecaes error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProduct := mock.NewMockUsecaseProduct(ctrl)

		p := &Product{
			cfg:            config.Config{},
			usecaseProduct: usecaseProduct,
		}

		reqUsecase := goproduct.ReqProductDetail{ID: 23423}
		resUsecase := goproduct.ResProductDetail{}

		usecaseProduct.EXPECT().
			GetDetail(context.Background(), reqUsecase).
			Return(resUsecase, assert.AnError)

		reqGRPC := &goproductgrpc.ReqProductDetail{Id: reqUsecase.ID}

		res, err := p.GetDetail(context.Background(), reqGRPC)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
