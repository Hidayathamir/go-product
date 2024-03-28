package grpc

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// RunServer run grpc server.
func RunServer(cfg config.Config, db *db.Postgres, rdb *redis.Client) error {
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
