package ctrlgrpc

import (
	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/db"
	"github.com/Hidayathamir/go-product/internal/usecase"
	"github.com/redis/go-redis/v9"
)

func injectionProduct(cfg config.Config, db *db.Postgres, rdb *redis.Client) *Product {
	repoProductCache := repopostgres.NewProductCache(cfg, rdb)
	repoProduct := repopostgres.NewProduct(cfg, db, repoProductCache)
	usecaseProduct := usecase.NewProduct(cfg, repoProduct)
	controllerProduct := newProduct(cfg, usecaseProduct)
	return controllerProduct
}

func injectionStock(cfg config.Config, db *db.Postgres) *Stock {
	repoStock := repopostgres.NewStock(cfg, db)
	usecaseStock := usecase.NewStock(cfg, repoStock)
	controllerStock := newStock(cfg, usecaseStock)
	return controllerStock
}