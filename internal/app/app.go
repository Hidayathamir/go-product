// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/dbpostgres"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseCLIArgs()

	cfg := initConfig(arg)

	handleCommandLineArgsMigrate(cfg, arg)

	newDBPostgres(cfg)
}

func newDBPostgres(cfg config.Config) *dbpostgres.Postgres {
	db, err := dbpostgres.NewPoolConn(cfg)
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}
	return db
}
