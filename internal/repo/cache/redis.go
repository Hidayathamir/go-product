package cache

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// Redis -.
type Redis struct {
	client *redis.Client
}

// NewRedis -.
func NewRedis(cfg config.Config) (*Redis, error) {
	addr := net.JoinHostPort(cfg.Redis.Host, strconv.Itoa(cfg.Redis.Port))

	var redisClient *redis.Client
	var err error
	for i := 0; i < 10; i++ {
		redisClient = redis.NewClient(&redis.Options{Addr: addr})
		err = redisClient.Ping(context.Background()).Err()
		if err != nil {
			err := fmt.Errorf("error ping redis: %w", err)
			logrus.
				WithField("attempt count", i+1).
				Warn(trace.Wrap(err))

			time.Sleep(time.Second)

			continue
		}
		break
	}

	if err != nil {
		err := fmt.Errorf("error 10 times when try to connect to redis: %w", err)
		return nil, trace.Wrap(err)
	}

	logrus.Info("success create redis connection 🟢")

	redis := &Redis{client: redisClient}

	return redis, nil
}
