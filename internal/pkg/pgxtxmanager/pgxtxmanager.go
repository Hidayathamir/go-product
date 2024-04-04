// Package pgxtxmanager -.
package pgxtxmanager

import (
	"context"
	"fmt"

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
			return fmt.Errorf("DBTx.Begin: %w", err)
		}
		ctx = context.WithValue(ctx, CtxKey, tx)
	}

	err := fn(ctx)

	if !isHasExternalTransaction {
		if err != nil {
			errRollback := tx.Rollback(ctx)
			if errRollback != nil {
				logrus.Warnf("pgx.Tx.Rollback: %v", errRollback)
			}
			return err
		}
		errCommit := tx.Commit(ctx)
		if errCommit != nil {
			return fmt.Errorf("pgx.Tx.Commit: %w", errCommit)
		}
	}

	return err
}

// GetTxFromContext -.
func GetTxFromContext(ctx context.Context) (pgx.Tx, bool) { //nolint:ireturn
	tx, ok := ctx.Value(CtxKey).(pgx.Tx)
	return tx, ok
}
