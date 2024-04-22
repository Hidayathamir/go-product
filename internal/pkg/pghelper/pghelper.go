// Package pghelper -.
package pghelper

import (
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/jackc/pgx/v5"
)

// GetPort return port from database URL.
func GetPort(dbURL string) (int, error) {
	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return 0, trace.Wrap(err)
	}
	return int(connConfig.Port), nil
}
