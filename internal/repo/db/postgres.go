package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// IPgxPool use this as dependency instead of *pgxpool.Pool so we can mock for
// unit test. Add method when needed.
type IPgxPool interface {
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

var _ IPgxPool = &pgxpool.Pool{}

// Postgres -.
type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pool    IPgxPool // use IPgxPool instead *pgxpool.Pool
}

// NewPGPoolConn return postgres pool connection.
func NewPGPoolConn(cfg config.Config) (*Postgres, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.PG.Username, cfg.PG.Password, cfg.PG.Host, cfg.PG.Port, cfg.PG.DBName,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	poolConfig.MaxConns = int32(cfg.PG.PoolMax)

	pg := &Postgres{
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	for i := 0; i < 10; i++ {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			err := fmt.Errorf("error create new conn pool: %w", err)
			logrus.
				WithField("attempt count", i+1).
				Warn(trace.Wrap(err))

			time.Sleep(time.Second)

			continue
		}

		err = pg.Pool.Ping(context.Background())
		if err != nil {
			err := fmt.Errorf("error ping db: %w", err)
			logrus.
				WithField("attempt count", i+1).
				Warn(trace.Wrap(err))

			time.Sleep(time.Second)

			continue
		}

		break
	}

	if err != nil {
		err := fmt.Errorf("error 10 times when try to create conn pool: %w", err)
		return nil, trace.Wrap(err)
	}

	logrus.Info("success create db connection pool 🟢")

	return pg, nil
}
