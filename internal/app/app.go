// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/sqldb/repopostgres/db"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseCLIArgs()

	cfg := initConfig(arg)

	handleCommandLineArgsMigrate(cfg, arg)

	newDBPostgres(cfg)
}

func newDBPostgres(cfg config.Config) *db.Postgres {
	db, err := db.NewPoolConn(cfg)
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}
	return db
}
