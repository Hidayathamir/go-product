package app

import (
	"path/filepath"

	"github.com/Hidayathamir/go-product/config"
	"github.com/Hidayathamir/go-product/internal/repo/repopostgres/db"
	"github.com/sirupsen/logrus"
)

// handleCommandLineArgsMigrate do db migration.
func handleCommandLineArgsMigrate(cfg config.Config, arg cliArg) {
	if arg.isIncludeMigrate {
		schemaMigrationPath := filepath.Join("internal", "repo", "repopostgres", "db", "schema_migration")
		err := db.MigrateUp(cfg, schemaMigrationPath)
		if err != nil {
			logrus.Fatalf("db.MigrateUp: %v", err)
		}
	}
}
