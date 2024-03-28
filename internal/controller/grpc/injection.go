package grpc

import (
	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/redis/go-redis/v9"
)

func injectionProduct(cfg config.Config, db *db.Postgres, rdb *redis.Client) *Product {
	repoProductCache := cache.NewProductCache(cfg, rdb)
	repoProduct := repo.NewProduct(cfg, db, repoProductCache)
	usecaseProduct := usecase.NewProduct(cfg, repoProduct)
	controllerProduct := newProduct(cfg, usecaseProduct)
	return controllerProduct
}

func injectionStock(cfg config.Config, db *db.Postgres) *Stock {
	repoStock := repo.NewStock(cfg, db)
	usecaseStock := usecase.NewStock(cfg, repoStock)
	controllerStock := newStock(cfg, usecaseStock)
	return controllerStock
}
