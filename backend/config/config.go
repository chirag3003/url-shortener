package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config holds all application configuration parsed from environment variables.
type Config struct {
	// Server
	Port string `env:"PORT" envDefault:"5000"`

	// PostgreSQL
	DatabaseURL string `env:"DATABASE_URL,required"`

	// Redis
	RedisURL string `env:"REDIS_URL,required"`

	// JWT
	JWTSecret     string        `env:"JWT_SECRET,required"`
	JWTExpiration time.Duration `env:"JWT_EXPIRATION" envDefault:"24h"`

	// AWS S3
	S3AccessKey string `env:"S3_ACCESS_KEY,required"`
	S3SecretKey string `env:"S3_SECRET_KEY,required"`
	S3Region    string `env:"S3_REGION,required"`
	S3Bucket    string `env:"S3_BUCKET,required"`
	S3Endpoint  string `env:"S3_ENDPOINT,required"`
	S3Folder    string `env:"S3_FOLDER" envDefault:"images"`

	// CORS
	CORSAllowOrigins string `env:"CORS_ALLOW_ORIGINS" envDefault:"*"`
	BaseURL          string `env:"BASE_URL" envDefault:"http://localhost:5000"`

	// Logging
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// Hyperflake
	HyperflakeDatacenterID int   `env:"HYPERFLAKE_DATACENTER_ID" envDefault:"1"`
	HyperflakeMachineID    int   `env:"HYPERFLAKE_MACHINE_ID" envDefault:"1"`
	HyperflakeEpochMS      int64 `env:"HYPERFLAKE_EPOCH_MS" envDefault:"0"`
}

// Load parses environment variables into a Config struct and validates required fields.
func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	if cfg.HyperflakeDatacenterID < 0 || cfg.HyperflakeDatacenterID > 31 {
		return nil, fmt.Errorf("HYPERFLAKE_DATACENTER_ID must be between 0 and 31")
	}
	if cfg.HyperflakeMachineID < 0 || cfg.HyperflakeMachineID > 31 {
		return nil, fmt.Errorf("HYPERFLAKE_MACHINE_ID must be between 0 and 31")
	}
	if cfg.HyperflakeEpochMS < 0 {
		return nil, fmt.Errorf("HYPERFLAKE_EPOCH_MS must be >= 0")
	}
	return cfg, nil
}
