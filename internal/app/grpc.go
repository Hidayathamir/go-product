package app

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	controllerHRPC "github.com/Hidayathamir/go-product/internal/controller/grpc"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// RunGRPCServer run grpc server.
func RunGRPCServer(cfg config.Config, db *db.Postgres, rdb *redis.Client) error {
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
	cProduct := injectionProduct(cfg, db, rdb)
	cStock := injectionStock(cfg, db)

	goproductgrpc.RegisterProductServer(grpcServer, cProduct)
	goproductgrpc.RegisterStockServer(grpcServer, cStock)
}

func injectionProduct(cfg config.Config, db *db.Postgres, rdb *redis.Client) *controllerHRPC.Product {
	repoProductCache := cache.NewProductCache(cfg, rdb)
	repoProduct := repo.NewProduct(cfg, db, repoProductCache)
	usecaseProduct := usecase.NewProduct(cfg, repoProduct)
	controllerProduct := controllerHRPC.NewProduct(cfg, usecaseProduct)
	return controllerProduct
}

func injectionStock(cfg config.Config, db *db.Postgres) *controllerHRPC.Stock {
	repoStock := repo.NewStock(cfg, db)
	usecaseStock := usecase.NewStock(cfg, repoStock)
	controllerStock := controllerHRPC.NewStock(cfg, usecaseStock)
	return controllerStock
}
