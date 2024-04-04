// Package pghelper -.
package pghelper

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

// GetPort return port from database URL.
func GetPort(dbURL string) (int, error) {
	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return 0, fmt.Errorf("pgx.ParseConfig: %w", err)
	}
	return int(connConfig.Port), nil
}
