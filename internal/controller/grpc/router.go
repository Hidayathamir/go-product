package grpc

import (
	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/Hidayathamir/go-product/pkg/goproductgrpc"
	"google.golang.org/grpc"
)

// This file contains all available servers.

func registerServer(cfg config.Config, grpcServer *grpc.Server, db *db.Postgres) {
	cProduct := injectionProduct(cfg, db)

	goproductgrpc.RegisterProductServer(grpcServer, cProduct)
}
