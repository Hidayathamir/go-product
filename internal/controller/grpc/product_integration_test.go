package controllergrpc

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestIntegrationProductSearch(t *testing.T) {
	t.Parallel()

	t.Run("search get 1 product", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
		require.NoError(t, err)

		redis := cache.NewRedis(cfg)

		repoProduct := repo.NewProduct(cfg, pg, redis)
		usecaseProduct := usecase.NewProduct(cfg, repoProduct)
		controllerProduct := NewProduct(cfg, usecaseProduct)

		p1Desc := "product 1 description"
		timeNow := time.Now()
		p1 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: p1Desc,
			Stock:       123,
		}
		p2 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Stock:       23,
		}
		p3 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Stock:       432,
		}
		p1ID, err := repoProduct.Add(context.Background(), p1)
		require.NoError(t, err)
		_, err = repoProduct.Add(context.Background(), p2)
		require.NoError(t, err)
		_, err = repoProduct.Add(context.Background(), p3)
		require.NoError(t, err)

		expectedRes := &goproductgrpc.ResProductSearch{
			Products: []*goproductgrpc.ResProductDetail{
				{
					Id:          p1ID,
					Sku:         p1.SKU,
					Slug:        p1.Slug,
					Name:        p1.Name,
					Description: p1.Description,
					Stock:       p1.Stock,
					CreatedAt:   timestamppb.New(timeNow),
					UpdatedAt:   timestamppb.New(timeNow),
				},
			},
		}

		req := &goproductgrpc.ReqProductSearch{
			Keyword: "uct 1 descrip",
		}
		res, err := controllerProduct.Search(context.Background(), req)

		require.NoError(t, err)
		require.Equal(t, len(expectedRes.GetProducts()), len(res.GetProducts()))
		assert.Equal(t, expectedRes.GetProducts()[0].GetId(), res.GetProducts()[0].GetId())
		assert.Equal(t, expectedRes.GetProducts()[0].GetSku(), res.GetProducts()[0].GetSku())
		assert.Equal(t, expectedRes.GetProducts()[0].GetSlug(), res.GetProducts()[0].GetSlug())
		assert.Equal(t, expectedRes.GetProducts()[0].GetName(), res.GetProducts()[0].GetName())
		assert.Equal(t, expectedRes.GetProducts()[0].GetDescription(), res.GetProducts()[0].GetDescription())
		assert.Equal(t, expectedRes.GetProducts()[0].GetStock(), res.GetProducts()[0].GetStock())
	})
	t.Run("search get many product", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
		require.NoError(t, err)

		redis := cache.NewRedis(cfg)

		repoProduct := repo.NewProduct(cfg, pg, redis)
		usecaseProduct := usecase.NewProduct(cfg, repoProduct)
		controllerProduct := NewProduct(cfg, usecaseProduct)

		p1Desc := "product 1 description"
		timeNow := time.Now()
		p1 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: p1Desc,
			Stock:       123,
		}
		p2 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: p1Desc,
			Stock:       23,
		}
		p3 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Stock:       432,
		}
		p1ID, err := repoProduct.Add(context.Background(), p1)
		require.NoError(t, err)
		p2ID, err := repoProduct.Add(context.Background(), p2)
		require.NoError(t, err)
		_, err = repoProduct.Add(context.Background(), p3)
		require.NoError(t, err)

		expectedRes := &goproductgrpc.ResProductSearch{
			Products: []*goproductgrpc.ResProductDetail{
				{
					Id:          p1ID,
					Sku:         p1.SKU,
					Slug:        p1.Slug,
					Name:        p1.Name,
					Description: p1.Description,
					Stock:       p1.Stock,
					CreatedAt:   timestamppb.New(timeNow),
					UpdatedAt:   timestamppb.New(timeNow),
				},
				{
					Id:          p2ID,
					Sku:         p2.SKU,
					Slug:        p2.Slug,
					Name:        p2.Name,
					Description: p2.Description,
					Stock:       p2.Stock,
					CreatedAt:   timestamppb.New(timeNow),
					UpdatedAt:   timestamppb.New(timeNow),
				},
			},
		}

		req := &goproductgrpc.ReqProductSearch{
			Keyword: "uct 1 descrip",
		}
		res, err := controllerProduct.Search(context.Background(), req)

		require.NoError(t, err)
		require.Equal(t, len(expectedRes.GetProducts()), len(res.GetProducts()))
		assert.Equal(t, expectedRes.GetProducts()[0].GetId(), res.GetProducts()[0].GetId())
		assert.Equal(t, expectedRes.GetProducts()[0].GetSku(), res.GetProducts()[0].GetSku())
		assert.Equal(t, expectedRes.GetProducts()[0].GetSlug(), res.GetProducts()[0].GetSlug())
		assert.Equal(t, expectedRes.GetProducts()[0].GetName(), res.GetProducts()[0].GetName())
		assert.Equal(t, expectedRes.GetProducts()[0].GetDescription(), res.GetProducts()[0].GetDescription())
		assert.Equal(t, expectedRes.GetProducts()[0].GetStock(), res.GetProducts()[0].GetStock())

		assert.Equal(t, expectedRes.GetProducts()[1].GetId(), res.GetProducts()[1].GetId())
		assert.Equal(t, expectedRes.GetProducts()[1].GetSku(), res.GetProducts()[1].GetSku())
		assert.Equal(t, expectedRes.GetProducts()[1].GetSlug(), res.GetProducts()[1].GetSlug())
		assert.Equal(t, expectedRes.GetProducts()[1].GetName(), res.GetProducts()[1].GetName())
		assert.Equal(t, expectedRes.GetProducts()[1].GetDescription(), res.GetProducts()[1].GetDescription())
		assert.Equal(t, expectedRes.GetProducts()[1].GetStock(), res.GetProducts()[1].GetStock())
	})
	t.Run("search get 0 product", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
		require.NoError(t, err)

		redis := cache.NewRedis(cfg)

		repoProduct := repo.NewProduct(cfg, pg, redis)
		usecaseProduct := usecase.NewProduct(cfg, repoProduct)
		controllerProduct := NewProduct(cfg, usecaseProduct)

		p1 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Stock:       123,
		}
		p2 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Stock:       23,
		}
		p3 := goproductdto.ReqProductAdd{
			SKU:         uuid.NewString(),
			Slug:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Stock:       432,
		}
		_, err = repoProduct.Add(context.Background(), p1)
		require.NoError(t, err)
		_, err = repoProduct.Add(context.Background(), p2)
		require.NoError(t, err)
		_, err = repoProduct.Add(context.Background(), p3)
		require.NoError(t, err)

		expectedRes := &goproductgrpc.ResProductSearch{
			Products: []*goproductgrpc.ResProductDetail{},
		}

		req := &goproductgrpc.ReqProductSearch{Keyword: uuid.NewString()}
		res, err := controllerProduct.Search(context.Background(), req)

		require.NoError(t, err)
		require.Equal(t, len(expectedRes.GetProducts()), len(res.GetProducts()))
	})
}
