package grpc

import (
	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/db"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

// This file contains all available servers.

func registerServer(cfg config.Config, grpcServer *grpc.Server, db *db.Postgres, rdb *redis.Client) {
	cProduct := injectionProduct(cfg, db, rdb)
	cStock := injectionStock(cfg, db)

	goproductgrpc.RegisterProductServer(grpcServer, cProduct)
	goproductgrpc.RegisterStockServer(grpcServer, cStock)
}
