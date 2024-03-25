package repo

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/pkg/query"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/repo/db/entity/table"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	sq "github.com/Masterminds/squirrel"
)

//go:generate mockgen -source=product.go -destination=mockrepo/product.go -package=mockrepo

// IProduct contains abstraction of repo product.
type IProduct interface {
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproduct.ResProductSearch, error)
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproduct.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproduct.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproduct.ResProductDetail, error)
}

// Product implement IProduct.
type Product struct {
	cfg config.Config
	db  *db.Postgres
}

var _ IProduct = &Product{}

// NewProduct return *Product which implement repo.IProduct.
func NewProduct(cfg config.Config, db *db.Postgres) *Product {
	return &Product{
		cfg: cfg,
		db:  db,
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
func (p *Product) GetDetailByID(context.Context, int64) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}

// GetDetailBySKU implements IProduct.
func (p *Product) GetDetailBySKU(context.Context, string) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}

// GetDetailBySlug implements IProduct.
func (p *Product) GetDetailBySlug(context.Context, string) (goproduct.ResProductDetail, error) {
	panic("unimplemented")
}
