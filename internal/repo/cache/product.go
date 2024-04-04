package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Hidayathamir/go-product/pkg/goproductdto"
)

//go:generate mockgen -source=product.go -destination=mockcache/product.go -package=mockcache

// IProduct contains abstraction of repo product cache.
type IProduct interface {
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproductdto.ResProductDetail, error)
	// SetDetailByID set product detail cache by id.
	SetDetailByID(ctx context.Context, data goproductdto.ResProductDetail, expire time.Duration) error

	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproductdto.ResProductDetail, error)
	// SetDetailBySKU get product detail cache by sku.
	SetDetailBySKU(ctx context.Context, data goproductdto.ResProductDetail, expire time.Duration) error

	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproductdto.ResProductDetail, error)
	// SetDetailBySlug get product detail cache by slug.
	SetDetailBySlug(ctx context.Context, data goproductdto.ResProductDetail, expire time.Duration) error
}

var _ IProduct = &Redis{}

///////////////////////////////// redis cache key /////////////////////////////////

const productRedisKeyPrefix = "product"

func (p *Redis) keyDetailByID(id int64) string {
	keyList := []string{p.cfg.App.Name, productRedisKeyPrefix, "DetailByID", strconv.FormatInt(id, 10)}
	return strings.Join(keyList, ":")
}

func (p *Redis) keyDetailBySKU(sku string) string {
	keyList := []string{p.cfg.App.Name, productRedisKeyPrefix, "DetailBySKU", sku}
	return strings.Join(keyList, ":")
}

func (p *Redis) keyDetailBySlug(slug string) string {
	keyList := []string{p.cfg.App.Name, productRedisKeyPrefix, "DetailBySlug", slug}
	return strings.Join(keyList, ":")
}

///////////////////////////////// redis cache key /////////////////////////////////

// GetDetailByID implements IProduct.
func (p *Redis) GetDetailByID(ctx context.Context, id int64) (goproductdto.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, p.keyDetailByID(id)).Result()
	if err != nil {
		return goproductdto.ResProductDetail{}, fmt.Errorf("Redis.rdb.Get: %w", err)
	}

	product := goproductdto.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		err := fmt.Errorf("json.Unmarshal, able to get value from redis but error when json unmarshal, trying to delete redis cache key: %w", err)

		errDel := p.rdb.Del(ctx, p.keyDetailByID(id)).Err()
		if errDel != nil {
			err = fmt.Errorf("Redis.rdb.Del, error delete redis cache key: %w", errDel)
		}

		return goproductdto.ResProductDetail{}, err
	}

	return product, nil
}

// SetDetailByID implements IProduct.
func (p *Redis) SetDetailByID(ctx context.Context, product goproductdto.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, p.keyDetailByID(product.ID), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("Redis.rdb.Set: %w", err)
	}

	return nil
}

// GetDetailBySKU implements IProduct.
func (p *Redis) GetDetailBySKU(ctx context.Context, sku string) (goproductdto.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, p.keyDetailBySKU(sku)).Result()
	if err != nil {
		return goproductdto.ResProductDetail{}, fmt.Errorf("Redis.rdb.Get: %w", err)
	}

	product := goproductdto.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		err := fmt.Errorf("json.Unmarshal, able to get value from redis but error when json unmarshal, trying to delete redis cache key: %w", err)

		errDel := p.rdb.Del(ctx, p.keyDetailBySKU(sku)).Err()
		if errDel != nil {
			err = fmt.Errorf("Redis.rdb.Del, error delete redis cache key: %w", errDel)
		}

		return goproductdto.ResProductDetail{}, err
	}

	return product, nil
}

// SetDetailBySKU implements IProduct.
func (p *Redis) SetDetailBySKU(ctx context.Context, product goproductdto.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, p.keyDetailBySKU(product.SKU), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("Redis.rdb.Set: %w", err)
	}

	return nil
}

// GetDetailBySlug implements IProduct.
func (p *Redis) GetDetailBySlug(ctx context.Context, slug string) (goproductdto.ResProductDetail, error) {
	val, err := p.rdb.Get(ctx, p.keyDetailBySlug(slug)).Result()
	if err != nil {
		return goproductdto.ResProductDetail{}, fmt.Errorf("Redis.rdb.Get: %w", err)
	}

	product := goproductdto.ResProductDetail{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		err := fmt.Errorf("json.Unmarshal, able to get value from redis but error when json unmarshal, trying to delete redis cache key: %w", err)

		errDel := p.rdb.Del(ctx, p.keyDetailBySlug(slug)).Err()
		if errDel != nil {
			err = fmt.Errorf("Redis.rdb.Del, error delete redis cache key: %w", errDel)
		}

		return goproductdto.ResProductDetail{}, err
	}

	return product, nil
}

// SetDetailBySlug implements IProduct.
func (p *Redis) SetDetailBySlug(ctx context.Context, product goproductdto.ResProductDetail, expire time.Duration) error {
	jsonByte, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = p.rdb.Set(ctx, p.keyDetailBySlug(product.Slug), string(jsonByte), expire).Err()
	if err != nil {
		return fmt.Errorf("Redis.rdb.Set: %w", err)
	}

	return nil
}
