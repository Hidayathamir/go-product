package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces/mock"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitProductSearch(t *testing.T) {
	t.Parallel()

	t.Run("search product success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductSearch{Keyword: "iphone12"}
		expectedRes := goproduct.ResProductSearch{
			Products: []goproduct.ResProductDetail{
				{
					ID:          23423,
					SKU:         "asdf",
					Slug:        "fesfes",
					Name:        "asfesf",
					Description: "sefsd",
					Stock:       234,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					ID:          32432,
					SKU:         "asdfes",
					Slug:        "fesfs",
					Name:        "fesfs",
					Description: "efsefq",
					Stock:       23422,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
		}

		repoProduct.EXPECT().
			Search(context.Background(), req.Keyword).
			Return(expectedRes, nil)

		res, err := p.Search(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
	t.Run("repo search error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductSearch{Keyword: "iphone12"}
		expectedRes := goproduct.ResProductSearch{}

		repoProduct.EXPECT().
			Search(context.Background(), req.Keyword).
			Return(expectedRes, assert.AnError)

		res, err := p.Search(context.Background(), req)

		assert.Equal(t, expectedRes, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("request validate error should return error", func(t *testing.T) {
		t.Parallel()

		t.Run("keyword empty should return error", func(t *testing.T) {
			t.Parallel()

			p := &Product{
				cfg: config.Config{},
			}

			req := goproduct.ReqProductSearch{Keyword: ""}
			expectedRes := goproduct.ResProductSearch{}

			res, err := p.Search(context.Background(), req)

			assert.Equal(t, expectedRes, res)
			require.Error(t, err)
			require.ErrorIs(t, err, goproduct.ErrRequestInvalid)
		})
	})
}

func TestUnitProductGetDetail(t *testing.T) {
	t.Parallel()

	t.Run("get detail by id success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductDetail{
			ID: 234,
		}
		expectedRes := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "asfes",
			Slug:        "gesges",
			Name:        "sefes",
			Description: "sfesf",
			Stock:       3522,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		repoProduct.EXPECT().
			GetDetailByID(context.Background(), req.ID).
			Return(expectedRes, nil)

		res, err := p.GetDetail(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
	t.Run("get detail by id error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductDetail{
			ID: 234,
		}
		expectedRes := goproduct.ResProductDetail{}

		repoProduct.EXPECT().
			GetDetailByID(context.Background(), req.ID).
			Return(expectedRes, assert.AnError)

		res, err := p.GetDetail(context.Background(), req)

		assert.Equal(t, expectedRes, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("get detail by sku success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductDetail{
			SKU: "asfes",
		}
		expectedRes := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "asfes",
			Slug:        "gesges",
			Name:        "sefes",
			Description: "sfesf",
			Stock:       3522,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		repoProduct.EXPECT().
			GetDetailBySKU(context.Background(), req.SKU).
			Return(expectedRes, nil)

		res, err := p.GetDetail(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
	t.Run("get detail by sku error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductDetail{
			SKU: "asfes",
		}
		expectedRes := goproduct.ResProductDetail{}

		repoProduct.EXPECT().
			GetDetailBySKU(context.Background(), req.SKU).
			Return(expectedRes, assert.AnError)

		res, err := p.GetDetail(context.Background(), req)

		assert.Equal(t, expectedRes, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("get detail by slug success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductDetail{
			Slug: "gesges",
		}
		expectedRes := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "asfes",
			Slug:        "gesges",
			Name:        "sefes",
			Description: "sfesf",
			Stock:       3522,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		repoProduct.EXPECT().
			GetDetailBySlug(context.Background(), req.Slug).
			Return(expectedRes, nil)

		res, err := p.GetDetail(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
	t.Run("get detail by slug error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProduct := mock.NewMockRepoProduct(ctrl)

		p := &Product{
			cfg:         config.Config{},
			repoProduct: repoProduct,
		}

		req := goproduct.ReqProductDetail{
			Slug: "gesges",
		}
		expectedRes := goproduct.ResProductDetail{}

		repoProduct.EXPECT().
			GetDetailBySlug(context.Background(), req.Slug).
			Return(expectedRes, assert.AnError)

		res, err := p.GetDetail(context.Background(), req)

		assert.Equal(t, expectedRes, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("request validate error should return error", func(t *testing.T) {
		t.Parallel()

		t.Run("id and sku and slug empty should return error", func(t *testing.T) {
			t.Parallel()

			p := &Product{
				cfg: config.Config{},
			}

			req := goproduct.ReqProductDetail{}
			expectedRes := goproduct.ResProductDetail{}

			res, err := p.GetDetail(context.Background(), req)

			assert.Equal(t, expectedRes, res)
			require.Error(t, err)
			require.ErrorIs(t, err, goproduct.ErrRequestInvalid)
		})
	})
}
