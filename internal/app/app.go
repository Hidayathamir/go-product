// Package app contains the application starter.
package app

import (
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseCLIArgs()

	cfg := initConfig(arg)

	handleCommandLineArgsMigrate(cfg, arg)

	db := newDBPostgres(cfg)

	rdb := newRedis(cfg)

	err := runGRPCServer(cfg, db, rdb)
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

func newRedis(cfg config.Config) *redis.Client {
	addr := net.JoinHostPort(cfg.Redis.Host, strconv.Itoa(cfg.Redis.Port))
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return rdb
}
