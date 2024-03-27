package repopostgres

import (
	"context"
	"testing"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/dbpostgres"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitStockIncrementStock(t *testing.T) {
	t.Parallel()

	t.Run("increment stock success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		s := &Stock{
			cfg: config.Config{},
			db: &dbpostgres.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		productID := int64(231)

		mockpool.ExpectExec("UPDATE").WithArgs(productID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = s.IncrementStock(context.Background(), productID)

		require.NoError(t, err)
	})
	t.Run("row affect 0 should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		s := &Stock{
			cfg: config.Config{},
			db: &dbpostgres.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		productID := int64(231)

		mockpool.ExpectExec("UPDATE").WithArgs(productID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 0))

		err = s.IncrementStock(context.Background(), productID)

		require.Error(t, err)
		require.ErrorContains(t, err, "pgconn.CommandTag.RowsAffected == 0")
	})
	t.Run("exec error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		s := &Stock{
			cfg: config.Config{},
			db: &dbpostgres.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		productID := int64(231)

		mockpool.ExpectExec("UPDATE").WithArgs(productID).WillReturnError(assert.AnError)

		err = s.IncrementStock(context.Background(), productID)

		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}

func TestUnitStockDecrementStock(t *testing.T) {
	t.Parallel()

	t.Run("increment stock success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		s := &Stock{
			cfg: config.Config{},
			db: &dbpostgres.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		productID := int64(231)

		mockpool.ExpectExec("UPDATE").WithArgs(productID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = s.DecrementStock(context.Background(), productID)

		require.NoError(t, err)
	})
	t.Run("row affect 0 should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		s := &Stock{
			cfg: config.Config{},
			db: &dbpostgres.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		productID := int64(231)

		mockpool.ExpectExec("UPDATE").WithArgs(productID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 0))

		err = s.DecrementStock(context.Background(), productID)

		require.Error(t, err)
		require.ErrorContains(t, err, "pgconn.CommandTag.RowsAffected == 0")
	})
	t.Run("exec error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		s := &Stock{
			cfg: config.Config{},
			db: &dbpostgres.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		productID := int64(231)

		mockpool.ExpectExec("UPDATE").WithArgs(productID).WillReturnError(assert.AnError)

		err = s.DecrementStock(context.Background(), productID)

		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
