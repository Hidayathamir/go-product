package app

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/controller/controllergrpc"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func runGRPCServer(cfg config.Config, db *db.Postgres, rdb *redis.Client) error {
	grpcServer := grpc.NewServer()

	registerServer(cfg, grpcServer, db, rdb)

	addr := net.JoinHostPort(cfg.GRPC.Host, strconv.Itoa(cfg.GRPC.Port))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	logrus.WithField("address", addr).Info("run grpc server")
	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("grpc.Server.Serve: %w", err)
	}

	return nil
}

func registerServer(cfg config.Config, grpcServer *grpc.Server, db *db.Postgres, rdb *redis.Client) {
	cProduct := injectionProductGRPC(cfg, db, rdb)
	cStock := injectionStockGRPC(cfg, db)

	goproductgrpc.RegisterProductServer(grpcServer, cProduct)
	goproductgrpc.RegisterStockServer(grpcServer, cStock)
}

func injectionProductGRPC(cfg config.Config, db *db.Postgres, rdb *redis.Client) *controllergrpc.Product {
	repoProductCache := cache.NewProductCache(cfg, rdb)
	repoProduct := repo.NewProduct(cfg, db, repoProductCache)
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
