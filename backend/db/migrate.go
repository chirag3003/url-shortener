package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

// MigrationFS returns the embedded migration filesystem.
func MigrationFS() embed.FS {
	return migrationFS
}

// RunMigrations applies all pending migrations.
func RunMigrations(ctx context.Context, dbURL string) error {
	goose.SetBaseFS(migrationFS)
	goose.SetTableName("goose_db_version")

	sqlDB, err := sql.Open("pgx", dbURL)
	if err != nil {
		return fmt.Errorf("open migration db: %w", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("ping migration db: %w", err)
	}

	if err := goose.UpContext(ctx, sqlDB, "migrations"); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
