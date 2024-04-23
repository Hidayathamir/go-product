// Package app contains the application starter.
package app

import (
	"flag"
	"path/filepath"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/Hidayathamir/go-product/internal/repo/cache"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
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

	err = runGRPCServer(cfg, dbPostgres, cacheRedis)
	if err != nil {
		logrus.Fatal(trace.Wrap(err))
	}
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
