package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/redis/go-redis/v9"
)

const productRedisKeyPrefix = "product"

var (
	keyDetailByID = func(id int64) string {
		keyList := []string{productRedisKeyPrefix, "GetDetailByID", strconv.FormatInt(id, 10)}
		return strings.Join(keyList, ":")
	}
	keyDetailBySKU = func(sku string) string {
		keyList := []string{productRedisKeyPrefix, "GetDetailBySKU", sku}
		return strings.Join(keyList, ":")
	}
	keyDetailBySlug = func(slug string) string {
		keyList := []string{productRedisKeyPrefix, "GetDetailBySlug", slug}
		return strings.Join(keyList, ":")
	}
)

// ProductCache implement IProductCache.
type ProductCache struct {
	cfg config.Config
	rdb *redis.Client
}

var _ IProductCache = &ProductCache{}

// NewProductCache return *ProductCache which implement repo.IProductCache.
func NewProductCache(cfg config.Config, rdb *redis.Client) *ProductCache {
	return &ProductCache{
		cfg: cfg,
		rdb: rdb,
	}
}

// GetDetailByID implements IProductCache.
func (p *ProductCache) GetDetailByID(ctx context.Context, id int64) (goproduct.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, keyDetailByID(id)).Result()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("ProductCache.rdb.Get: %w", err)
	}

	product := goproduct.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return product, nil
}

// SetDetailByID implements IProductCache.
func (p *ProductCache) SetDetailByID(ctx context.Context, product goproduct.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, keyDetailByID(product.ID), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("ProductCache.rdb.Set: %w", err)
	}

	return nil
}

// GetDetailBySKU implements IProductCache.
func (p *ProductCache) GetDetailBySKU(ctx context.Context, sku string) (goproduct.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, keyDetailBySKU(sku)).Result()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("ProductCache.rdb.Get: %w", err)
	}

	product := goproduct.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return product, nil
}

// SetDetailBySKU implements IProductCache.
func (p *ProductCache) SetDetailBySKU(ctx context.Context, product goproduct.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, keyDetailBySKU(product.SKU), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("ProductCache.rdb.Set: %w", err)
	}

	return nil
}

// GetDetailBySlug implements IProductCache.
func (p *ProductCache) GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, keyDetailBySlug(slug)).Result()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("ProductCache.rdb.Get: %w", err)
	}

	product := goproduct.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return product, nil
}

// SetDetailBySlug implements IProductCache.
func (p *ProductCache) SetDetailBySlug(ctx context.Context, product goproduct.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, keyDetailBySlug(product.Slug), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("ProductCache.rdb.Set: %w", err)
	}

	return nil
}
