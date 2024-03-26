package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/pkg/query"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/repo/db/entity/table"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=product.go -destination=mockrepo/product.go -package=mockrepo

// IProduct contains abstraction of repo product.
type IProduct interface {
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error)
}

// Product implement IProduct.
type Product struct {
	cfg   config.Config
	db    *db.Postgres
	cache IProductCache
}

var _ IProduct = &Product{}

// NewProduct return *Product which implement repo.IProduct.
func NewProduct(cfg config.Config, db *db.Postgres, cache IProductCache) *Product {
	return &Product{
		cfg:   cfg,
		db:    db,
		cache: cache,
	}
}

// Search implements IProduct.
func (p *Product) Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error) {
	sql, args, err := p.db.Builder.
		Select(
			table.Product.Dot.ID, table.Product.Dot.SKU, table.Product.Dot.Slug,
			table.Product.Dot.Name, table.Product.Dot.Description, table.Stock.Dot.Stock,
			table.Product.Dot.CreatedAt, table.Product.Dot.UpdatedAt).
		From(table.Product.String()).
		Join(query.JoinOnEqual(table.Stock.String(), table.Product.Dot.ID, table.Stock.Dot.ProductID)).
		Where(sq.Or{
			sq.ILike{table.Product.Dot.Name: "%" + keyword + "%"},
			sq.ILike{table.Product.Dot.Description: "%" + keyword + "%"},
		}).
		ToSql()
	if err != nil {
		return goproduct.ResProductSearch{}, fmt.Errorf("Product.db.Builder.ToSql: %w", err)
	}

	rows, err := p.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return goproduct.ResProductSearch{}, fmt.Errorf("Product.db.Pool.Query: %w", err)
	}
	defer rows.Close()

	products := []goproduct.ResProductDetail{}
	for rows.Next() {
		product := goproduct.ResProductDetail{}

		err := rows.Scan(
			&product.ID, &product.SKU, &product.Slug, &product.Name,
			&product.Description, &product.Stock, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return goproduct.ResProductSearch{}, fmt.Errorf("pgx.Rows.Scan: %w", err)
		}

		products = append(products, product)
	}

	return goproduct.ResProductSearch{Products: products}, nil
}

// GetDetailByID implements IProduct.
func (p *Product) GetDetailByID(ctx context.Context, id int64) (goproduct.ResProductDetail, error) { //nolint:dupl
	product, err := p.cache.GetDetailByID(ctx, id)
	if err == nil {
		return product, nil
	}

	logrus.Warnf("Product.cache.GetDetailByID: %v", err)

	sql, args, err := p.db.Builder.
		Select(
			table.Product.Dot.ID, table.Product.Dot.SKU, table.Product.Dot.Slug,
			table.Product.Dot.Name, table.Product.Dot.Description, table.Stock.Dot.Stock,
			table.Product.Dot.CreatedAt, table.Product.Dot.UpdatedAt).
		From(table.Product.String()).
		Join(query.JoinOnEqual(table.Stock.String(), table.Product.Dot.ID, table.Stock.Dot.ProductID)).
		Where(sq.Eq{table.Product.Dot.ID: id}).
		ToSql()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("Product.db.Builder.ToSql: %w", err)
	}

	product = goproduct.ResProductDetail{}
	err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&product.ID, &product.SKU, &product.Slug,
		&product.Name, &product.Description, &product.Stock,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		err := fmt.Errorf("Product.db.Pool.QueryRow: %w", err)
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("%w: %w", goproduct.ErrProductNotFound, err)
		}
		return goproduct.ResProductDetail{}, err
	}

	err = p.cache.SetDetailByID(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		logrus.Warnf("Product.cache.SetDetailByID: %v", err)
	}

	return product, nil
}

// GetDetailBySKU implements IProduct.
func (p *Product) GetDetailBySKU(ctx context.Context, sku string) (goproduct.ResProductDetail, error) { //nolint:dupl
	product, err := p.cache.GetDetailBySKU(ctx, sku)
	if err == nil {
		return product, nil
	}

	logrus.Warnf("Product.cache.GetDetailBySKU: %v", err)

	sql, args, err := p.db.Builder.
		Select(
			table.Product.Dot.ID, table.Product.Dot.SKU, table.Product.Dot.Slug,
			table.Product.Dot.Name, table.Product.Dot.Description, table.Stock.Dot.Stock,
			table.Product.Dot.CreatedAt, table.Product.Dot.UpdatedAt).
		From(table.Product.String()).
		Join(query.JoinOnEqual(table.Stock.String(), table.Product.Dot.ID, table.Stock.Dot.ProductID)).
		Where(sq.Eq{table.Product.Dot.SKU: sku}).
		ToSql()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("Product.db.Builder.ToSql: %w", err)
	}

	product = goproduct.ResProductDetail{}
	err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&product.ID, &product.SKU, &product.Slug,
		&product.Name, &product.Description, &product.Stock,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		err := fmt.Errorf("Product.db.Pool.QueryRow: %w", err)
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("%w: %w", goproduct.ErrProductNotFound, err)
		}
		return goproduct.ResProductDetail{}, err
	}

	err = p.cache.SetDetailBySKU(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		logrus.Warnf("Product.cache.SetDetailBySKU: %v", err)
	}

	return product, nil
}

// GetDetailBySlug implements IProduct.
func (p *Product) GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error) { //nolint:dupl
	product, err := p.cache.GetDetailBySlug(ctx, slug)
	if err == nil {
		return product, nil
	}

	logrus.Warnf("Product.cache.GetDetailBySlug: %v", err)

	sql, args, err := p.db.Builder.
		Select(
			table.Product.Dot.ID, table.Product.Dot.SKU, table.Product.Dot.Slug,
			table.Product.Dot.Name, table.Product.Dot.Description, table.Stock.Dot.Stock,
			table.Product.Dot.CreatedAt, table.Product.Dot.UpdatedAt).
		From(table.Product.String()).
		Join(query.JoinOnEqual(table.Stock.String(), table.Product.Dot.ID, table.Stock.Dot.ProductID)).
		Where(sq.Eq{table.Product.Dot.Slug: slug}).
		ToSql()
	if err != nil {
		return goproduct.ResProductDetail{}, fmt.Errorf("Product.db.Builder.ToSql: %w", err)
	}

	product = goproduct.ResProductDetail{}
	err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&product.ID, &product.SKU, &product.Slug,
		&product.Name, &product.Description, &product.Stock,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		err := fmt.Errorf("Product.db.Pool.QueryRow: %w", err)
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("%w: %w", goproduct.ErrProductNotFound, err)
		}
		return goproduct.ResProductDetail{}, err
	}

	err = p.cache.SetDetailBySlug(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		logrus.Warnf("Product.cache.SetDetailBySlug: %v", err)
	}

	return product, nil
}
