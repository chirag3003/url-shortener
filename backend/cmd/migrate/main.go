package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"database/sql"
	"github.com/chirag3003/go-backend-template/db"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_URL is required")
		os.Exit(1)
	}

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sqlDB, err := sql.Open("pgx", databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open db: %v\n", err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	if err := sqlDB.PingContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "ping db: %v\n", err)
		os.Exit(1)
	}

	goose.SetBaseFS(db.MigrationFS())
	goose.SetTableName("goose_db_version")

	switch command {
	case "up":
		err = goose.UpContext(ctx, sqlDB, "migrations")
	case "down":
		err = goose.DownContext(ctx, sqlDB, "migrations")
	case "status":
		err = goose.StatusContext(ctx, sqlDB, "migrations")
	default:
		fmt.Fprintf(os.Stderr, "unsupported command %q (use: up, down, status)\n", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "migration command failed: %v\n", err)
		os.Exit(1)
	}
}
