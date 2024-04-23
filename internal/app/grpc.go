package app

import (
	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	transportgrpc "github.com/Hidayathamir/go-product/internal/transport/grpc"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"google.golang.org/grpc"
)

func registerGRPCServer(cfg config.Config, grpcServer *grpc.Server, dbPostgres *db.Postgres, cacheRedis *cache.Redis) {
	tProduct := injectionProductGRPC(cfg, dbPostgres, cacheRedis)
	tStock := injectionStockGRPC(cfg, dbPostgres)

	goproductgrpc.RegisterProductServer(grpcServer, tProduct)
	goproductgrpc.RegisterStockServer(grpcServer, tStock)
}

func injectionProductGRPC(cfg config.Config, dbPostgres *db.Postgres, cacheRedis *cache.Redis) *transportgrpc.Product {
	cacheProduct := cache.NewProduct(cfg, cacheRedis)
	repoProduct := repo.NewProduct(cfg, dbPostgres, cacheProduct)
	usecaseProduct := usecase.NewProduct(cfg, repoProduct)
	transportProduct := transportgrpc.NewProduct(cfg, usecaseProduct)
	return transportProduct
}

func injectionStockGRPC(cfg config.Config, dbPostgres *db.Postgres) *transportgrpc.Stock {
	repoStock := repo.NewStock(cfg, dbPostgres)
	usecaseStock := usecase.NewStock(cfg, repoStock)
	transportStock := transportgrpc.NewStock(cfg, usecaseStock)
	return transportStock
}
