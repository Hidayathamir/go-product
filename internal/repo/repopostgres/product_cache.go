package repopostgres

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/redis/go-redis/v9"
)

// ProductCache implement IProductCache.
type ProductCache struct {
	cfg config.Config
	rdb *redis.Client
}

var _ interfaces.RepoProductCache = &ProductCache{}

// NewProductCache return *ProductCache which implement repo.IProductCache.
func NewProductCache(cfg config.Config, rdb *redis.Client) *ProductCache {
	return &ProductCache{
		cfg: cfg,
		rdb: rdb,
	}
}

///////////////////////////////// redis cache key /////////////////////////////////

const productRedisKeyPrefix = "product"

func (p *ProductCache) keyDetailByID(id int64) string {
	keyList := []string{p.cfg.App.Name, productRedisKeyPrefix, "DetailByID", strconv.FormatInt(id, 10)}
	return strings.Join(keyList, ":")
}

func (p *ProductCache) keyDetailBySKU(sku string) string {
	keyList := []string{p.cfg.App.Name, productRedisKeyPrefix, "DetailBySKU", sku}
	return strings.Join(keyList, ":")
}

func (p *ProductCache) keyDetailBySlug(slug string) string {
	keyList := []string{p.cfg.App.Name, productRedisKeyPrefix, "DetailBySlug", slug}
	return strings.Join(keyList, ":")
}

///////////////////////////////// redis cache key /////////////////////////////////

// GetDetailByID implements IProductCache.
func (p *ProductCache) GetDetailByID(ctx context.Context, id int64) (goproduct.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, p.keyDetailByID(id)).Result()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("ProductCache.rdb.Get: %w", err)
	}

	product := goproduct.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		err := fmt.Errorf("json.Unmarshal, able to get value from redis but error when json unmarshal, trying to delete redis cache key: %w", err)

		errDel := p.rdb.Del(ctx, p.keyDetailByID(id)).Err()
		if errDel != nil {
			err = fmt.Errorf("ProductCache.rdb.Del, error delete redis cache key: %w", errDel)
		}

		return goproduct.ResProductDetail{}, err
	}

	return product, nil
}

// SetDetailByID implements IProductCache.
func (p *ProductCache) SetDetailByID(ctx context.Context, product goproduct.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, p.keyDetailByID(product.ID), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("ProductCache.rdb.Set: %w", err)
	}

	return nil
}

// GetDetailBySKU implements IProductCache.
func (p *ProductCache) GetDetailBySKU(ctx context.Context, sku string) (goproduct.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, p.keyDetailBySKU(sku)).Result()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("ProductCache.rdb.Get: %w", err)
	}

	product := goproduct.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		err := fmt.Errorf("json.Unmarshal, able to get value from redis but error when json unmarshal, trying to delete redis cache key: %w", err)

		errDel := p.rdb.Del(ctx, p.keyDetailBySKU(sku)).Err()
		if errDel != nil {
			err = fmt.Errorf("ProductCache.rdb.Del, error delete redis cache key: %w", errDel)
		}

		return goproduct.ResProductDetail{}, err
	}

	return product, nil
}

// SetDetailBySKU implements IProductCache.
func (p *ProductCache) SetDetailBySKU(ctx context.Context, product goproduct.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, p.keyDetailBySKU(product.SKU), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("ProductCache.rdb.Set: %w", err)
	}

	return nil
}

// GetDetailBySlug implements IProductCache.
func (p *ProductCache) GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, p.keyDetailBySlug(slug)).Result()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("ProductCache.rdb.Get: %w", err)
	}

	product := goproduct.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		err := fmt.Errorf("json.Unmarshal, able to get value from redis but error when json unmarshal, trying to delete redis cache key: %w", err)

		errDel := p.rdb.Del(ctx, p.keyDetailBySlug(slug)).Err()
		if errDel != nil {
			err = fmt.Errorf("ProductCache.rdb.Del, error delete redis cache key: %w", errDel)
		}

		return goproduct.ResProductDetail{}, err
	}

	return product, nil
}

// SetDetailBySlug implements IProductCache.
func (p *ProductCache) SetDetailBySlug(ctx context.Context, product goproduct.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, p.keyDetailBySlug(product.Slug), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("ProductCache.rdb.Set: %w", err)
	}

	return nil
}
