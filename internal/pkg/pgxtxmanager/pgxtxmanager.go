// Package pgxtxmanager -.
package pgxtxmanager

import (
	"context"

	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// DBTx -.
type DBTx interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type ctxKey string

// CtxKey -.
var CtxKey = ctxKey("pgxtxmanager-sql-transaction")

// SQLTransaction -.
func SQLTransaction(ctx context.Context, dbTx DBTx, fn func(context.Context) error) error {
	tx, isHasExternalTransaction := ctx.Value(CtxKey).(pgx.Tx)

	if !isHasExternalTransaction {
		var err error
		tx, err = dbTx.Begin(ctx)
		if err != nil {
			return trace.Wrap(err)
		}
		ctx = context.WithValue(ctx, CtxKey, tx)
	}

	err := fn(ctx)

	if !isHasExternalTransaction {
		if err != nil {
			errRollback := tx.Rollback(ctx)
			if errRollback != nil {
				logrus.Warn(trace.Wrap(errRollback))
			}
			return err
		}
		errCommit := tx.Commit(ctx)
		if errCommit != nil {
			return trace.Wrap(errCommit)
		}
	}

	return err
}

// GetTxFromContext -.
func GetTxFromContext(ctx context.Context) (pgx.Tx, bool) { //nolint:ireturn
	tx, ok := ctx.Value(CtxKey).(pgx.Tx)
	return tx, ok
}
