package grpc

import (
	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/internal/usecase"
)

func injectionProduct(cfg config.Config, db *db.Postgres) *Product {
	repoProduct := repo.NewProduct(cfg, db)
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
