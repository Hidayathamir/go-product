package app

import (
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	controllergrpc "github.com/Hidayathamir/go-product/internal/transport/grpc"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func runGRPCServer(cfg config.Config, db *db.Postgres, redis *cache.Redis) error {
	grpcServer := grpc.NewServer()

	registerServer(cfg, grpcServer, db, redis)

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

func registerServer(cfg config.Config, grpcServer *grpc.Server, db *db.Postgres, redis *cache.Redis) {
	cProduct := injectionProductGRPC(cfg, db, redis)
	cStock := injectionStockGRPC(cfg, db)

	goproductgrpc.RegisterProductServer(grpcServer, cProduct)
	goproductgrpc.RegisterStockServer(grpcServer, cStock)
}

func injectionProductGRPC(cfg config.Config, db *db.Postgres, redis *cache.Redis) *controllergrpc.Product {
	cache := cache.NewProduct(cfg, redis)
	repoProduct := repo.NewProduct(cfg, db, cache)
	usecaseProduct := usecase.NewProduct(cfg, repoProduct)
	controllerProduct := controllergrpc.NewProduct(cfg, usecaseProduct)
	return controllerProduct
}

func injectionStockGRPC(cfg config.Config, db *db.Postgres) *controllergrpc.Stock {
	repoStock := repo.NewStock(cfg, db)
	usecaseStock := usecase.NewStock(cfg, repoStock)
	controllerStock := controllergrpc.NewStock(cfg, usecaseStock)
	return controllerStock
}
