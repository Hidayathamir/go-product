// Package app contains the application starter.
package app

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run application.
func Run() { //nolint:funlen
	arg := parseCLIArgs()

	cfg := initConfig(arg)

	handleCLIArgs(cfg, arg)

	dbPostgres, err := db.NewPGPoolConn(cfg)
	if err != nil {
		logrus.Fatal(trace.Wrap(err))
	}

	cacheRedis, err := cache.NewRedis(cfg)
	if err != nil {
		logrus.Fatal(trace.Wrap(err))
	}

	logrus.Info("initializing grpc server in a goroutine so that it won't block the graceful shutdown handling below")
	var grpcServer *grpc.Server
	go func() {
		grpcServer = grpc.NewServer()

		registerGRPCServer(cfg, grpcServer, dbPostgres, cacheRedis)

		addr := net.JoinHostPort(cfg.GRPC.Host, strconv.Itoa(cfg.GRPC.Port))
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			logrus.Fatal(trace.Wrap(err))
		}

		logrus.WithField("address", addr).Info("run grpc server")
		err = grpcServer.Serve(lis)
		if err != nil {
			logrus.Fatal(trace.Wrap(err))
		}
	}()

	logrus.Info("initializing http server in a goroutine so that it won't block the graceful shutdown handling below")
	var httpServer *http.Server
	go func() {
		registerHTTPRouter()

		addr := net.JoinHostPort(cfg.HTTP.Host, strconv.Itoa(cfg.HTTP.Port))
		httpServer = &http.Server{ //nolint:gosec
			Addr:    addr,
			Handler: nil, // TODO: ADD GIN AS ROUTER
		}

		logrus.WithField("address", addr).Info("run http server")
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatal(trace.Wrap(err))
		}
	}()

	logrus.Info("listens for the interrupt signal from the OS")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logrus.Info("listen for the interrupt signal")
	<-ctx.Done()

	stop()
	logrus.Info("shutting down gracefully, press Ctrl+C again to force")

	var gracefulShutdownWG sync.WaitGroup

	logrus.Info("shutting down gracefully grpc server")
	gracefulShutdownWG.Add(1)
	go func() {
		grpcServer.GracefulStop()
		gracefulShutdownWG.Done()
	}()

	logrus.Info("shutting down gracefully http server")
	gracefulShutdownWG.Add(1)
	go func() {
		logrus.Info("inform http server it has 10 seconds to finish the request it is currently handling")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //nolint:gomnd
		defer cancel()
		err = httpServer.Shutdown(ctx)
		if err != nil {
			logrus.Fatal(trace.Wrap(err))
		}
		gracefulShutdownWG.Done()
	}()

	logrus.Info("wait graceful shutdown finish")
	gracefulShutdownWG.Wait()
	logrus.Info("graceful shutdown finish")
}

type cliArg struct {
	isIncludeMigrate bool
	isLoadEnv        bool
}

func parseCLIArgs() cliArg {
	arg := cliArg{}

	flag.BoolVar(&arg.isIncludeMigrate, "include-migrate", false, "is include migrate, if true will do migrate before run app, default false.")
	flag.BoolVar(&arg.isLoadEnv, "load-env", false, "is load env var, if true load env var and override config, default false.")

	flag.Usage = cleanenv.FUsage(flag.CommandLine.Output(), &config.Config{}, nil, flag.Usage)

	flag.Parse()

	return arg
}

func initConfig(arg cliArg) config.Config {
	yamlPath := filepath.Join("internal", "config", "config.yml")

	var cfgLoader config.Loader
	if arg.isLoadEnv {
		cfgLoader = &config.EnvLoader{YAMLPath: yamlPath}
	} else {
		cfgLoader = &config.YamlLoader{Path: yamlPath}
	}

	cfg, err := config.Init(cfgLoader)
	if err != nil {
		logrus.Fatal(trace.Wrap(err))
	}

	return cfg
}

func handleCLIArgs(cfg config.Config, arg cliArg) {
	if arg.isIncludeMigrate {
		schemaMigrationPath := filepath.Join("internal", "repo", "db", "schema_migration")
		err := db.MigrateUp(cfg, schemaMigrationPath)
		if err != nil {
			logrus.Fatal(trace.Wrap(err))
		}
	}
}
