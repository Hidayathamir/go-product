package repo

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/repo/db/entity/table"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitProductSearch(t *testing.T) {
	t.Parallel()

	t.Run("search product success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		p1 := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "sku1",
			Slug:        "slug1",
			Name:        "name1",
			Description: "desc1",
			Stock:       123,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		p2 := goproduct.ResProductDetail{
			ID:          23,
			SKU:         "sku2",
			Slug:        "slug2",
			Name:        "name2",
			Description: "desc2",
			Stock:       323,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		expectedProducts := []goproduct.ResProductDetail{p1, p2}
		expectedRes := goproduct.ResProductSearch{Products: expectedProducts}

		keyword := "iphone"
		keywordIlike := "%" + keyword + "%"

		mockpool.ExpectQuery("SELECT").WithArgs(keywordIlike, keywordIlike).
			WillReturnRows(
				pgxmock.NewRows([]string{
					table.Product.ID, table.Product.SKU, table.Product.Slug,
					table.Product.Name, table.Product.Description, table.Stock.Stock,
					table.Product.CreatedAt, table.Product.UpdatedAt,
				}).
					AddRow(p1.ID, p1.SKU, p1.Slug, p1.Name, p1.Description, p1.Stock, p1.CreatedAt, p1.UpdatedAt).
					AddRow(p2.ID, p2.SKU, p2.Slug, p2.Name, p2.Description, p2.Stock, p2.CreatedAt, p2.UpdatedAt),
			)

		res, err := p.Search(context.Background(), keyword)

		require.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
	t.Run("query error should return error", func(t *testing.T) { t.Parallel() }) // TODO: IMPLEMENT
	t.Run("scan error should return error", func(t *testing.T) { t.Parallel() })
	t.Run("sql builder error should return error", func(t *testing.T) { t.Parallel() })
}
