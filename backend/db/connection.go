package db

import (
	"context"
	"fmt"
	"time"

	"github.com/chirag3003/go-backend-template/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// Connection provides access to PostgreSQL and cleanup.
type Connection interface {
	Close()
	Pool() *pgxpool.Pool
	Ping(ctx context.Context) error
}

type conn struct {
	pool *pgxpool.Pool
	log  zerolog.Logger
}

func (c *conn) Close() {
	c.pool.Close()
	c.log.Info().Msg("disconnected from PostgreSQL")
}

func (c *conn) Pool() *pgxpool.Pool {
	return c.pool
}

func (c *conn) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

// ConnectPostgres establishes a connection pool to PostgreSQL.
func ConnectPostgres(cfg *config.Config, log zerolog.Logger) (Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Database URL", cfg.DatabaseURL)
	poolCfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing database url: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("connecting to PostgreSQL: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pinging PostgreSQL: %w", err)
	}

	log.Info().Msg("connected to PostgreSQL")

	return &conn{
		pool: pool,
		log:  log,
	}, nil
}
