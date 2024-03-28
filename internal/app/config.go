package app

import (
	"path/filepath"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/sirupsen/logrus"
)

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
		logrus.Fatalf("config.Init: %v", err)
	}

	return cfg
}
