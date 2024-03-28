package repopostgres

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/sqldb/repopostgres/db"
	"github.com/Hidayathamir/go-product/internal/repo/sqldb/repopostgres/db/table"
	"github.com/Hidayathamir/go-product/internal/usecase/usecaseinterfaces/mock"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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
	t.Run("query error should return error", func(t *testing.T) {
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

		keyword := "iphone"
		keywordIlike := "%" + keyword + "%"

		mockpool.ExpectQuery("SELECT").WithArgs(keywordIlike, keywordIlike).
			WillReturnError(assert.AnError)

		res, err := p.Search(context.Background(), keyword)

		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("scan error should return error", func(t *testing.T) {
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

		keyword := "iphone"
		keywordIlike := "%" + keyword + "%"

		mockpool.ExpectQuery("SELECT").WithArgs(keywordIlike, keywordIlike).
			WillReturnRows(pgxmock.NewRows([]string{table.Product.ID}).AddRow(123))

		res, err := p.Search(context.Background(), keyword)

		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorContains(t, err, "pgx.Rows.Scan")
	})
}

func TestUnitProductGetDetailByID(t *testing.T) {
	t.Parallel()

	t.Run("get detail success with cache", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		p1 := goproduct.ResProductDetail{
			ID:          2141,
			SKU:         "sku",
			Slug:        "slug",
			Name:        "name",
			Description: "desc",
			Stock:       23,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		cache.EXPECT().
			GetDetailByID(context.Background(), p1.ID).
			Return(p1, nil)

		res, err := p.GetDetailByID(context.Background(), p1.ID)
		require.NoError(t, err)
		assert.Equal(t, p1, res)
	})
	t.Run("get detail success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		p1 := goproduct.ResProductDetail{
			ID:          2141,
			SKU:         "sku",
			Slug:        "slug",
			Name:        "name",
			Description: "desc",
			Stock:       23,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		cache.EXPECT().
			GetDetailByID(context.Background(), p1.ID).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(p1.ID).
			WillReturnRows(
				pgxmock.NewRows([]string{
					table.Product.ID, table.Product.SKU, table.Product.Slug,
					table.Product.Name, table.Product.Description, table.Stock.Stock,
					table.Product.CreatedAt, table.Product.UpdatedAt,
				}).AddRow(p1.ID, p1.SKU, p1.Slug, p1.Name, p1.Description, p1.Stock, p1.CreatedAt, p1.UpdatedAt),
			)

		cache.EXPECT().
			SetDetailByID(context.Background(), p1, goproduct.DefaultCacheExpire).
			Return(nil)

		res, err := p.GetDetailByID(context.Background(), p1.ID)
		require.NoError(t, err)
		assert.Equal(t, p1, res)
	})
	t.Run("query error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		id := int64(2342)

		cache.EXPECT().
			GetDetailByID(context.Background(), id).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(id).WillReturnError(assert.AnError)

		res, err := p.GetDetailByID(context.Background(), id)
		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("query error no rows should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		id := int64(2342)

		cache.EXPECT().
			GetDetailByID(context.Background(), id).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(id).WillReturnError(pgx.ErrNoRows)

		res, err := p.GetDetailByID(context.Background(), id)
		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, goproduct.ErrProductNotFound)
	})
}

func TestUnitProductGetDetailBySKU(t *testing.T) {
	t.Parallel()

	t.Run("get detail success with cache", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		p1 := goproduct.ResProductDetail{
			ID:          2141,
			SKU:         "sku",
			Slug:        "slug",
			Name:        "name",
			Description: "desc",
			Stock:       23,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		cache.EXPECT().
			GetDetailBySKU(context.Background(), p1.SKU).
			Return(p1, nil)

		res, err := p.GetDetailBySKU(context.Background(), p1.SKU)
		require.NoError(t, err)
		assert.Equal(t, p1, res)
	})
	t.Run("get detail success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		p1 := goproduct.ResProductDetail{
			ID:          2141,
			SKU:         "sku",
			Slug:        "slug",
			Name:        "name",
			Description: "desc",
			Stock:       23,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		cache.EXPECT().
			GetDetailBySKU(context.Background(), p1.SKU).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(p1.SKU).
			WillReturnRows(
				pgxmock.NewRows([]string{
					table.Product.ID, table.Product.SKU, table.Product.Slug,
					table.Product.Name, table.Product.Description, table.Stock.Stock,
					table.Product.CreatedAt, table.Product.UpdatedAt,
				}).AddRow(p1.ID, p1.SKU, p1.Slug, p1.Name, p1.Description, p1.Stock, p1.CreatedAt, p1.UpdatedAt),
			)

		cache.EXPECT().
			SetDetailBySKU(context.Background(), p1, goproduct.DefaultCacheExpire).
			Return(nil)

		res, err := p.GetDetailBySKU(context.Background(), p1.SKU)
		require.NoError(t, err)
		assert.Equal(t, p1, res)
	})
	t.Run("query error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		sku := "skutest"

		cache.EXPECT().
			GetDetailBySKU(context.Background(), sku).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(sku).WillReturnError(assert.AnError)

		res, err := p.GetDetailBySKU(context.Background(), sku)
		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("query error no rows should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		sku := "skutest"

		cache.EXPECT().
			GetDetailBySKU(context.Background(), sku).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(sku).WillReturnError(pgx.ErrNoRows)

		res, err := p.GetDetailBySKU(context.Background(), sku)
		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, goproduct.ErrProductNotFound)
	})
}

func TestUnitProductGetDetailBySlug(t *testing.T) {
	t.Parallel()

	t.Run("get detail success with cache", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		p1 := goproduct.ResProductDetail{
			ID:          2141,
			SKU:         "sku",
			Slug:        "slug",
			Name:        "name",
			Description: "desc",
			Stock:       23,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		cache.EXPECT().
			GetDetailBySlug(context.Background(), p1.Slug).
			Return(p1, nil)

		res, err := p.GetDetailBySlug(context.Background(), p1.Slug)
		require.NoError(t, err)
		assert.Equal(t, p1, res)
	})
	t.Run("get detail success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		p1 := goproduct.ResProductDetail{
			ID:          2141,
			SKU:         "sku",
			Slug:        "slug",
			Name:        "name",
			Description: "desc",
			Stock:       23,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		cache.EXPECT().
			GetDetailBySlug(context.Background(), p1.Slug).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(p1.Slug).
			WillReturnRows(
				pgxmock.NewRows([]string{
					table.Product.ID, table.Product.SKU, table.Product.Slug,
					table.Product.Name, table.Product.Description, table.Stock.Stock,
					table.Product.CreatedAt, table.Product.UpdatedAt,
				}).AddRow(p1.ID, p1.SKU, p1.Slug, p1.Name, p1.Description, p1.Stock, p1.CreatedAt, p1.UpdatedAt),
			)

		cache.EXPECT().
			SetDetailBySlug(context.Background(), p1, goproduct.DefaultCacheExpire).
			Return(nil)

		res, err := p.GetDetailBySlug(context.Background(), p1.Slug)
		require.NoError(t, err)
		assert.Equal(t, p1, res)
	})
	t.Run("query error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		slug := "slugtest"

		cache.EXPECT().
			GetDetailBySlug(context.Background(), slug).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(slug).WillReturnError(assert.AnError)

		res, err := p.GetDetailBySlug(context.Background(), slug)
		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("query error no rows should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cache := mock.NewMockRepoProductCache(ctrl)

		p := &Product{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
			cache: cache,
		}

		slug := "slugtest"

		cache.EXPECT().
			GetDetailBySlug(context.Background(), slug).
			Return(goproduct.ResProductDetail{}, assert.AnError)

		mockpool.ExpectQuery("SELECT").WithArgs(slug).WillReturnError(pgx.ErrNoRows)

		res, err := p.GetDetailBySlug(context.Background(), slug)
		assert.Empty(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, goproduct.ErrProductNotFound)
	})
}
