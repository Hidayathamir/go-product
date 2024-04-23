package controllergrpc

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/Hidayathamir/go-product/internal/config"
	"github.com/Hidayathamir/go-product/internal/pkg/pghelper"
	"github.com/Hidayathamir/go-product/internal/repo/db"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// initTestIntegration init config, create pg container, update config port
// based on pg container, do db migrations.
func initTestIntegration(t *testing.T) config.Config {
	t.Helper()

	cfg := configInit(t)

	pgContainer := createPGContainer(t, cfg)
	t.Cleanup(func() { require.NoError(t, pgContainer.Terminate(context.Background())) })

	updateConfigPGPort(t, &cfg, pgContainer)

	dbMigrateUp(t, cfg)

	s := miniredis.RunT(t)
	updateConfigRedis(t, &cfg, s)

	return cfg
}

func updateConfigRedis(t *testing.T, cfg *config.Config, s *miniredis.Miniredis) {
	t.Helper()

	cfg.Redis.Host = s.Host()

	port, err := strconv.Atoi(s.Port())
	require.NoError(t, err)

	cfg.Redis.Port = port
}

func configInit(t *testing.T) config.Config {
	t.Helper()

	require.NoError(t, os.Setenv("LOGGER_LOG_LEVEL", "fatal"))

	configYamlPath := filepath.Join("..", "..", "..", "internal", "config", "config.yml")
	cfg, err := config.Init(&config.EnvLoader{YAMLPath: configYamlPath})
	require.NoError(t, err)
	return cfg
}

type mute struct{}

func (n mute) Printf(string, ...interface{}) {}

// createPGContainer create pg container.
func createPGContainer(t *testing.T, cfg config.Config) *postgres.PostgresContainer {
	t.Helper()

	pgContainer, err := postgres.RunContainer(context.Background(),
		testcontainers.WithLogger(&mute{}),
		testcontainers.WithImage("postgres:16"),
		postgres.WithDatabase(cfg.PG.DBName),
		postgres.WithUsername(cfg.PG.Username),
		postgres.WithPassword(cfg.PG.Password),
		testcontainers.WithWaitStrategy(
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)
	return pgContainer
}

// updateConfigPGPort update config port based on pg container, do db migrations.
func updateConfigPGPort(t *testing.T, cfg *config.Config, pgContainer *postgres.PostgresContainer) {
	t.Helper()

	dbURL, err := pgContainer.ConnectionString(context.Background())
	require.NoError(t, err)

	port, err := pghelper.GetPort(dbURL)
	require.NoError(t, err)

	cfg.PG.Port = port
}

// dbMigrateUp do db migrations.
func dbMigrateUp(t *testing.T, cfg config.Config) {
	t.Helper()

	schemaMigrationPath := filepath.Join("..", "..", "..", "internal", "repo", "db", "schema_migration")
	require.NoError(t, db.MigrateUp(cfg, schemaMigrationPath))
}
