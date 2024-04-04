package cache

import (
	"net"
	"strconv"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/redis/go-redis/v9"
)

// Redis -.
type Redis struct {
	cfg config.Config
	rdb *redis.Client
}

// NewRedis -.
func NewRedis(cfg config.Config) *Redis {
	addr := net.JoinHostPort(cfg.Redis.Host, strconv.Itoa(cfg.Redis.Port))
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return &Redis{
		cfg: cfg,
		rdb: rdb,
	}
}
