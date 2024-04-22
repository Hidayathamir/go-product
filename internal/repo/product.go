package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/entity"
	"github.com/Hidayathamir/go-product/internal/pkg/pgxtxmanager"
	"github.com/Hidayathamir/go-product/internal/pkg/query"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/repo/db/table"
	"github.com/Hidayathamir/go-product/pkg/goproduct"
	"github.com/Hidayathamir/go-product/pkg/goproductdto"
	"github.com/Hidayathamir/go-product/pkg/goproducterror"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=product.go -destination=mockrepo/product.go -package=mockrepo

// IProduct contains abstraction of repo product.
type IProduct interface {
	// GetDetailByID get product detail by id.
	GetDetailByID(ctx context.Context, ID int64) (goproductdto.ResProductDetail, error)
	// GetDetailBySKU get product detail by sku.
	GetDetailBySKU(ctx context.Context, SKU string) (goproductdto.ResProductDetail, error)
	// GetDetailBySlug get product detail by slug.
	GetDetailBySlug(ctx context.Context, slug string) (goproductdto.ResProductDetail, error)
	// Search search product by name or description using keyword.
	Search(ctx context.Context, keyword string) (goproductdto.ResProductSearch, error)
	// Add adds product to database.
	Add(ctx context.Context, req goproductdto.ReqProductAdd) (int64, error)
}

// Product implement IProduct.
type Product struct {
	cfg   config.Config
	db    *db.Postgres
	cache cache.IProduct
}

var _ IProduct = &Product{}

// NewProduct return *Product which implement repo.IProduct.
func NewProduct(cfg config.Config, db *db.Postgres, cache cache.IProduct) *Product {
	return &Product{
		cfg:   cfg,
		db:    db,
		cache: cache,
	}
}

// Search implements IProduct.
func (p *Product) Search(ctx context.Context, keyword string) (goproductdto.ResProductSearch, error) {
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
		return goproductdto.ResProductSearch{}, trace.Wrap(err)
	}

	var rows pgx.Rows
	if tx, ok := pgxtxmanager.GetTxFromContext(ctx); ok {
		rows, err = tx.Query(ctx, sql, args...)
	} else {
		rows, err = p.db.Pool.Query(ctx, sql, args...)
	}
	if err != nil {
		return goproductdto.ResProductSearch{}, trace.Wrap(err)
	}
	defer rows.Close()

	products := []goproductdto.ResProductDetail{}
	for rows.Next() {
		product := goproductdto.ResProductDetail{}

		err := rows.Scan(
			&product.ID, &product.SKU, &product.Slug, &product.Name,
			&product.Description, &product.Stock, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return goproductdto.ResProductSearch{}, trace.Wrap(err)
		}

		products = append(products, product)
	}

	return goproductdto.ResProductSearch{Products: products}, nil
}

// GetDetailByID implements IProduct.
func (p *Product) GetDetailByID(ctx context.Context, id int64) (goproductdto.ResProductDetail, error) { //nolint:dupl
	product, err := p.cache.GetDetailByID(ctx, id)
	if err == nil {
		return product, nil
	}

	logrus.Warn(trace.Wrap(err))

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
		return goproductdto.ResProductDetail{}, trace.Wrap(err)
	}

	var row pgx.Row
	if tx, ok := pgxtxmanager.GetTxFromContext(ctx); ok {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = p.db.Pool.QueryRow(ctx, sql, args...)
	}

	product = goproductdto.ResProductDetail{}
	err = row.Scan(
		&product.ID, &product.SKU, &product.Slug,
		&product.Name, &product.Description, &product.Stock,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		err := trace.Wrap(err)
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("%w: %w", goproducterror.ErrProductNotFound, err)
		}
		return goproductdto.ResProductDetail{}, err
	}

	err = p.cache.SetDetailByID(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		logrus.Warn(trace.Wrap(err))
	}

	return product, nil
}

// GetDetailBySKU implements IProduct.
func (p *Product) GetDetailBySKU(ctx context.Context, sku string) (goproductdto.ResProductDetail, error) {
	product, err := p.cache.GetDetailBySKU(ctx, sku)
	if err == nil {
		return product, nil
	}

	logrus.Warn(trace.Wrap(err))

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
		return goproductdto.ResProductDetail{}, trace.Wrap(err)
	}

	var row pgx.Row
	if tx, ok := pgxtxmanager.GetTxFromContext(ctx); ok {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = p.db.Pool.QueryRow(ctx, sql, args...)
	}

	product = goproductdto.ResProductDetail{}
	err = row.Scan(
		&product.ID, &product.SKU, &product.Slug,
		&product.Name, &product.Description, &product.Stock,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		err := trace.Wrap(err)
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("%w: %w", goproducterror.ErrProductNotFound, err)
		}
		return goproductdto.ResProductDetail{}, err
	}

	err = p.cache.SetDetailBySKU(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		logrus.Warnf("Product.cache.SetDetailBySKU: %v", err)
	}

	return product, nil
}

// GetDetailBySlug implements IProduct.
func (p *Product) GetDetailBySlug(ctx context.Context, slug string) (goproductdto.ResProductDetail, error) { //nolint:dupl
	product, err := p.cache.GetDetailBySlug(ctx, slug)
	if err == nil {
		return product, nil
	}

	logrus.Warn(trace.Wrap(err))

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
		return goproductdto.ResProductDetail{}, trace.Wrap(err)
	}

	var row pgx.Row
	if tx, ok := pgxtxmanager.GetTxFromContext(ctx); ok {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = p.db.Pool.QueryRow(ctx, sql, args...)
	}

	product = goproductdto.ResProductDetail{}
	err = row.Scan(
		&product.ID, &product.SKU, &product.Slug,
		&product.Name, &product.Description, &product.Stock,
		&product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		err := trace.Wrap(err)
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("%w: %w", goproducterror.ErrProductNotFound, err)
		}
		return goproductdto.ResProductDetail{}, err
	}

	err = p.cache.SetDetailBySlug(ctx, product, goproduct.DefaultCacheExpire)
	if err != nil {
		logrus.Warn(trace.Wrap(err))
	}

	return product, nil
}

// Add implements IProduct.
func (p *Product) Add(ctx context.Context, req goproductdto.ReqProductAdd) (int64, error) {
	var productID int64
	err := pgxtxmanager.SQLTransaction(ctx, p.db.Pool, func(ctx context.Context) error {
		timeNow := time.Now()
		product := entity.Product{
			SKU:         req.SKU,
			Slug:        req.Slug,
			Name:        req.Name,
			Description: req.Description,
			CreatedAt:   timeNow,
			UpdatedAt:   timeNow,
		}

		var err error
		productID, err = p.addProduct(ctx, product)
		if err != nil {
			return trace.Wrap(err)
		}

		stock := entity.Stock{
			ProductID: productID,
			Stock:     req.Stock,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		}

		_, err = p.addStock(ctx, stock)
		if err != nil {
			return trace.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return 0, trace.Wrap(err)
	}

	return productID, nil
}

func (p *Product) addProduct(ctx context.Context, product entity.Product) (int64, error) {
	sql, args, err := p.db.Builder.
		Insert(table.Product.String()).
		Columns(
			table.Product.SKU, table.Product.Slug, table.Product.Name,
			table.Product.Description, table.Product.CreatedAt, table.Product.UpdatedAt,
		).
		Values(
			product.SKU, product.Slug, product.Name,
			product.Description, product.CreatedAt, product.UpdatedAt,
		).
		Suffix(query.Returning(table.Product.ID)).
		ToSql()
	if err != nil {
		return 0, trace.Wrap(err)
	}

	var row pgx.Row
	if tx, ok := pgxtxmanager.GetTxFromContext(ctx); ok {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = p.db.Pool.QueryRow(ctx, sql, args...)
	}

	var productID int64
	err = row.Scan(&productID)
	if err != nil {
		err := trace.Wrap(err)

		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			return 0, err
		}

		isErrDuplicateSKU := pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == table.Product.Constraint.ProductUnique
		if isErrDuplicateSKU {
			return 0, fmt.Errorf("%w: %w", goproducterror.ErrProductDuplicateSKU, err)
		}

		isErrDuplicateSlug := pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == table.Product.Constraint.ProductUnique1
		if isErrDuplicateSlug {
			return 0, fmt.Errorf("%w: %w", goproducterror.ErrProductDuplicateSlug, err)
		}
	}

	return productID, nil
}

func (p *Product) addStock(ctx context.Context, stock entity.Stock) (int64, error) {
	sql, args, err := p.db.Builder.
		Insert(table.Stock.String()).
		Columns(
			table.Stock.ProductID, table.Stock.Stock,
			table.Stock.CreatedAt, table.Stock.UpdatedAt,
		).
		Values(
			stock.ProductID, stock.Stock,
			stock.CreatedAt, stock.UpdatedAt,
		).
		Suffix(query.Returning(table.Stock.ID)).
		ToSql()
	if err != nil {
		return 0, trace.Wrap(err)
	}

	var row pgx.Row
	if tx, ok := pgxtxmanager.GetTxFromContext(ctx); ok {
		row = tx.QueryRow(ctx, sql, args...)
	} else {
		row = p.db.Pool.QueryRow(ctx, sql, args...)
	}

	var stockID int64
	err = row.Scan(&stockID)
	if err != nil {
		err := trace.Wrap(err)

		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			return 0, err
		}

		isErrDuplicateProductID := pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == table.Stock.Constraint.StockUnique
		if isErrDuplicateProductID {
			return 0, fmt.Errorf("%w: %w", goproducterror.ErrStockDuplicateProductID, err)
		}
	}

	return stockID, nil
}
