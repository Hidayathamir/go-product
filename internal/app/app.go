// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
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
		logrus.Fatal(trace.Wrap(err))
	}
}

func newDBPostgres(cfg config.Config) *db.Postgres {
	db, err := db.NewPGPoolConn(cfg)
	if err != nil {
		logrus.Fatal(trace.Wrap(err))
	}
	return db
}
