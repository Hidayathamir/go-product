package repopostgres

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/db"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/db/entity/table"
	"github.com/Hidayathamir/go-product/internal/usecase/interfaces"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// Stock implement IStock.
type Stock struct {
	cfg config.Config
	db  *db.Postgres
}

var _ interfaces.RepoStock = &Stock{}

// NewStock return *Stock which implement repo.IStock.
func NewStock(cfg config.Config, db *db.Postgres) *Stock {
	return &Stock{
		cfg: cfg,
		db:  db,
	}
}

// IncrementStock implements IStock.
func (s *Stock) IncrementStock(ctx context.Context, productID int64) error {
	sql, args, err := s.db.Builder.
		Update(table.Stock.String()).
		Set(table.Stock.Stock, sq.Expr(table.Stock.Stock+"+1")).
		Where(sq.Eq{
			table.Stock.ProductID: productID,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("Stock.db.Builder.ToSql: %w", err)
	}

	commandTag, err := s.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Stock.db.Pool.Exec: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("pgconn.CommandTag.RowsAffected == 0: %w", pgx.ErrNoRows)
	}

	return nil
}

// DecrementStock implements IStock.
func (s *Stock) DecrementStock(ctx context.Context, productID int64) error {
	sql, args, err := s.db.Builder.
		Update(table.Stock.String()).
		Set(table.Stock.Stock, sq.Expr(table.Stock.Stock+"-1")).
		Where(sq.Eq{
			table.Stock.ProductID: productID,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("Stock.db.Builder.ToSql: %w", err)
	}

	commandTag, err := s.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Stock.db.Pool.Exec: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("pgconn.CommandTag.RowsAffected == 0: %w", pgx.ErrNoRows)
	}

	return nil
}
