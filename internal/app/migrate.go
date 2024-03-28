package app

import (
	"path/filepath"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/repo/sqldb/repopostgres/db"
	"github.com/sirupsen/logrus"
)

// handleCommandLineArgsMigrate do db migration.
func handleCommandLineArgsMigrate(cfg config.Config, arg cliArg) {
	if arg.isIncludeMigrate {
		schemaMigrationPath := filepath.Join("internal", "repo", "repopostgres", "dbpostgres", "schema_migration")
		err := db.MigrateUp(cfg, schemaMigrationPath)
		if err != nil {
			logrus.Fatalf("db.MigrateUp: %v", err)
		}
	}
}
