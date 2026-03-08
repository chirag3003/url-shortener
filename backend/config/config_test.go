package config_test

import (
	"os"
	"testing"

	"github.com/chirag3003/go-backend-template/config"
)

func setTestEnv(t *testing.T) {
	t.Helper()
	envVars := map[string]string{
		"PORT":                     "3000",
		"DATABASE_URL":             "postgres://postgres:postgres@localhost:5432/url_shortener?sslmode=disable",
		"JWT_SECRET":               "test-secret",
		"JWT_EXPIRATION":           "2h",
		"S3_ACCESS_KEY":            "test-key",
		"S3_SECRET_KEY":            "test-secret-key",
		"S3_REGION":                "us-east-1",
		"S3_BUCKET":                "test-bucket",
		"S3_ENDPOINT":              "s3.amazonaws.com",
		"S3_FOLDER":                "uploads",
		"CORS_ALLOW_ORIGINS":       "http://localhost:3000",
		"BASE_URL":                 "http://localhost:5000",
		"LOG_LEVEL":                "debug",
		"HYPERFLAKE_DATACENTER_ID": "1",
		"HYPERFLAKE_MACHINE_ID":    "1",
		"HYPERFLAKE_EPOCH_MS":      "0",
	}

	for k, v := range envVars {
		t.Setenv(k, v)
	}
}

func TestLoad_Success(t *testing.T) {
	setTestEnv(t)

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.Port != "3000" {
		t.Errorf("expected Port '3000', got '%s'", cfg.Port)
	}
	if cfg.DatabaseURL != "postgres://postgres:postgres@localhost:5432/url_shortener?sslmode=disable" {
		t.Errorf("expected DatabaseURL, got '%s'", cfg.DatabaseURL)
	}
	if cfg.JWTSecret != "test-secret" {
		t.Errorf("expected JWTSecret 'test-secret', got '%s'", cfg.JWTSecret)
	}
	if cfg.S3Bucket != "test-bucket" {
		t.Errorf("expected S3Bucket 'test-bucket', got '%s'", cfg.S3Bucket)
	}
	if cfg.CORSAllowOrigins != "http://localhost:3000" {
		t.Errorf("expected CORSAllowOrigins, got '%s'", cfg.CORSAllowOrigins)
	}
}

func TestLoad_MissingRequired(t *testing.T) {
	// Clear all env vars that config requires
	os.Clearenv()

	_, err := config.Load()
	if err == nil {
		t.Fatal("Load should return error when required vars are missing")
	}
}

func TestLoad_Defaults(t *testing.T) {
	// Set only required variables
	required := map[string]string{
		"DATABASE_URL":  "postgres://postgres:postgres@localhost:5432/url_shortener?sslmode=disable",
		"JWT_SECRET":    "secret",
		"S3_ACCESS_KEY": "key",
		"S3_SECRET_KEY": "secret",
		"S3_REGION":     "us-east-1",
		"S3_BUCKET":     "bucket",
		"S3_ENDPOINT":   "endpoint",
	}

	for k, v := range required {
		t.Setenv(k, v)
	}

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	// Check defaults
	if cfg.Port != "5000" {
		t.Errorf("expected default Port '5000', got '%s'", cfg.Port)
	}
	if cfg.S3Folder != "images" {
		t.Errorf("expected default S3Folder 'images', got '%s'", cfg.S3Folder)
	}
	if cfg.CORSAllowOrigins != "*" {
		t.Errorf("expected default CORSAllowOrigins '*', got '%s'", cfg.CORSAllowOrigins)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("expected default LogLevel 'info', got '%s'", cfg.LogLevel)
	}
}
