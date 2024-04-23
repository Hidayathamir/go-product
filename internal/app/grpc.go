package app

import (
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	transportgrpc "github.com/Hidayathamir/go-product/internal/transport/grpc"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func runGRPCServer(cfg config.Config, dbPostgres *db.Postgres, cacheRedis *cache.Redis) error {
	grpcServer := grpc.NewServer()

	registerServer(cfg, grpcServer, dbPostgres, cacheRedis)

	addr := net.JoinHostPort(cfg.GRPC.Host, strconv.Itoa(cfg.GRPC.Port))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return trace.Wrap(err)
	}

	logrus.WithField("address", addr).Info("run grpc server")
	err = grpcServer.Serve(lis)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

func registerServer(cfg config.Config, grpcServer *grpc.Server, dbPostgres *db.Postgres, cacheRedis *cache.Redis) {
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
