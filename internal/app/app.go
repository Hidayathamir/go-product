// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseCLIArgs()

	cfg := initConfig(arg)

	handleCommandLineArgsMigrate(cfg, arg)

	db := newDBPostgres(cfg)

	redis := cache.NewRedis(cfg)

	err := runGRPCServer(cfg, db, redis)
	if err != nil {
		logrus.Fatalf("runGRPCServer: %v", err)
	}
}

func newDBPostgres(cfg config.Config) *db.Postgres {
	db, err := db.NewPGPoolConn(cfg)
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}
	return db
}
