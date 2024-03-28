package reporedis

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/jutil"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitProductCacheGetDetailByID(t *testing.T) {
	t.Parallel()

	t.Run("get detail by id success", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		expectedProduct := goproduct.ResProductDetail{
			ID:          234231,
			SKU:         "asdf",
			Slug:        "aefsf",
			Name:        "fsdfea",
			Description: "asdfes",
			Stock:       234,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		err := s.Set(p.keyDetailByID(expectedProduct.ID), jutil.ToJSONString(expectedProduct))
		require.NoError(t, err)

		product, err := p.GetDetailByID(context.Background(), expectedProduct.ID)

		require.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
	})
	t.Run("key not found should return error", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		product, err := p.GetDetailByID(context.Background(), 23423)

		assert.Empty(t, product)
		require.Error(t, err)
		require.ErrorIs(t, err, redis.Nil)
	})
	t.Run("json marshal error should return error", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		id := int64(23423)
		err := s.Set(p.keyDetailByID(id), "plain string will error json unmarshal to struct")
		require.NoError(t, err)

		product, err := p.GetDetailByID(context.Background(), id)

		assert.Empty(t, product)
		require.Error(t, err)
		require.ErrorContains(t, err, "json.Unmarshal")
	})
}

func TestUnitProductCacheSetDetailByID(t *testing.T) {
	t.Parallel()

	t.Run("set detail by id success", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		product := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "asdfads",
			Slug:        "fese",
			Name:        "fssda",
			Description: "aesfs",
			Stock:       234,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		expire := time.Hour

		err := p.SetDetailByID(context.Background(), product, expire)
		require.NoError(t, err)

		t.Run("after set to redis, should able to get, and the value is equal with expected", func(t *testing.T) {
			product2, err := p.GetDetailByID(context.Background(), product.ID)
			require.NoError(t, err)

			assert.Equal(t, product, product2)
		})
		s.FastForward(expire + time.Second)
		t.Run("after expire, should not able to get", func(t *testing.T) {
			product2, err := p.GetDetailByID(context.Background(), product.ID)

			assert.Empty(t, product2)
			require.Error(t, err)
			require.ErrorIs(t, err, redis.Nil)
		})
	})
}

func TestUnitProductCacheGetDetailBySKU(t *testing.T) {
	t.Parallel()

	t.Run("get detail by sku success", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		expectedProduct := goproduct.ResProductDetail{
			ID:          234231,
			SKU:         "asdf",
			Slug:        "aefsf",
			Name:        "fsdfea",
			Description: "asdfes",
			Stock:       234,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		err := s.Set(p.keyDetailBySKU(expectedProduct.SKU), jutil.ToJSONString(expectedProduct))
		require.NoError(t, err)

		product, err := p.GetDetailBySKU(context.Background(), expectedProduct.SKU)

		require.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
	})
	t.Run("key not found should return error", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		product, err := p.GetDetailBySKU(context.Background(), "dummy sku")

		assert.Empty(t, product)
		require.Error(t, err)
		require.ErrorIs(t, err, redis.Nil)
	})
	t.Run("json marshal error should return error", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		sku := "dummysku"
		err := s.Set(p.keyDetailBySKU(sku), "plain string will error json unmarshal to struct")
		require.NoError(t, err)

		product, err := p.GetDetailBySKU(context.Background(), sku)

		assert.Empty(t, product)
		require.Error(t, err)
		require.ErrorContains(t, err, "json.Unmarshal")
	})
}

func TestUnitProductCacheSetDetailBySKU(t *testing.T) {
	t.Parallel()

	t.Run("set detail by sku success", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		product := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "asdfads",
			Slug:        "fese",
			Name:        "fssda",
			Description: "aesfs",
			Stock:       234,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		expire := time.Hour

		err := p.SetDetailBySKU(context.Background(), product, expire)
		require.NoError(t, err)

		t.Run("after set to redis, should able to get, and the value is equal with expected", func(t *testing.T) {
			product2, err := p.GetDetailBySKU(context.Background(), product.SKU)
			require.NoError(t, err)

			assert.Equal(t, product, product2)
		})
		s.FastForward(expire + time.Second)
		t.Run("after expire, should not able to get", func(t *testing.T) {
			product2, err := p.GetDetailBySKU(context.Background(), product.SKU)

			assert.Empty(t, product2)
			require.Error(t, err)
			require.ErrorIs(t, err, redis.Nil)
		})
	})
}

func TestUnitProductCacheGetDetailBySlug(t *testing.T) {
	t.Parallel()

	t.Run("get detail by sku success", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		expectedProduct := goproduct.ResProductDetail{
			ID:          234231,
			SKU:         "asdf",
			Slug:        "aefsf",
			Name:        "fsdfea",
			Description: "asdfes",
			Stock:       234,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		err := s.Set(p.keyDetailBySlug(expectedProduct.Slug), jutil.ToJSONString(expectedProduct))
		require.NoError(t, err)

		product, err := p.GetDetailBySlug(context.Background(), expectedProduct.Slug)

		require.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
	})
	t.Run("key not found should return error", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		slug := "jlkdfsd"

		product, err := p.GetDetailBySlug(context.Background(), slug)

		assert.Empty(t, product)
		require.Error(t, err)
		require.ErrorIs(t, err, redis.Nil)
	})
	t.Run("json marshal error should return error", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		slug := "asdfes"
		err := s.Set(p.keyDetailBySlug(slug), "plain string will error json unmarshal to struct")
		require.NoError(t, err)

		product, err := p.GetDetailBySlug(context.Background(), slug)

		assert.Empty(t, product)
		require.Error(t, err)
		require.ErrorContains(t, err, "json.Unmarshal")
	})
}

func TestUnitProductCacheSetDetailBySlug(t *testing.T) {
	t.Parallel()

	t.Run("set detail by sku success", func(t *testing.T) {
		t.Parallel()

		s := miniredis.RunT(t)

		rdb := redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		})

		p := &Product{
			cfg: config.Config{},
			rdb: rdb,
		}

		product := goproduct.ResProductDetail{
			ID:          234,
			SKU:         "asdfads",
			Slug:        "fese",
			Name:        "fssda",
			Description: "aesfs",
			Stock:       234,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		expire := time.Hour

		err := p.SetDetailBySlug(context.Background(), product, expire)
		require.NoError(t, err)

		t.Run("after set to redis, should able to get, and the value is equal with expected", func(t *testing.T) {
			product2, err := p.GetDetailBySlug(context.Background(), product.Slug)
			require.NoError(t, err)

			assert.Equal(t, product, product2)
		})
		s.FastForward(expire + time.Second)
		t.Run("after expire, should not able to get", func(t *testing.T) {
			product2, err := p.GetDetailBySlug(context.Background(), product.Slug)

			assert.Empty(t, product2)
			require.Error(t, err)
			require.ErrorIs(t, err, redis.Nil)
		})
	})
}
