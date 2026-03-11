package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/chirag3003/go-backend-template/config"
	"github.com/chirag3003/go-backend-template/controller"
	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/helpers/aws"
	"github.com/chirag3003/go-backend-template/middleware"
	"github.com/chirag3003/go-backend-template/pkg/auth"
	"github.com/chirag3003/go-backend-template/pkg/cache"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/pkg/logger"
	"github.com/chirag3003/go-backend-template/pkg/messaging"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/chirag3003/go-backend-template/routes"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
	"github.com/joho/godotenv"
	redisclient "github.com/redis/go-redis/v9"
)

func main() {
	// Load .env file (optional in production)
	_ = godotenv.Load()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log := logger.New(cfg.LogLevel)

	// Connect to PostgreSQL
	dbConn, err := db.ConnectPostgres(cfg, log)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to PostgreSQL")
	}
	defer dbConn.Close()

	// Run database migrations
	if err := db.RunMigrations(context.Background(), cfg.DatabaseURL); err != nil {
		log.Fatal().Err(err).Msg("failed to run database migrations")
	}

	// Initialize Redis
	if err := cache.Init(cfg.RedisURL); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize Redis")
	}
	defer cache.GetDefault().Close()

	// Setup AWS
	if err := aws.SetupAWS(cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to setup AWS")
	}

	// Initialize services
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTExpiration)
	idgen.Init(
		cfg.HyperflakeDatacenterID,
		cfg.HyperflakeMachineID,
		cfg.HyperflakeEpochMS,
	)

	// Initialize repositories
	repo := repository.NewRepository(dbConn)

	// Initialize services
	authService := service.NewAuthService(repo.User, jwtService, log)
	userService := service.NewUserService(repo.User, log)
	mediaService := service.NewMediaService(repo.Media, repo.S3, cfg, log)
	linkService := service.NewLinkService(repo.Link, cfg.BaseURL, log)
	analyticsService := service.NewAnalyticsService(repo.Click, repo.Link, log)

	// Start Analytics Worker
	opts, err := redisclient.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse redis URL")
	}
	rClient := redisclient.NewClient(opts)
	streamManager := messaging.NewStreamManager(rClient)
	streamManager.EnsureGroup(context.Background(), messaging.AnalyticsStream, messaging.AnalyticsGroup)
	go streamManager.Consume(context.Background(), messaging.AnalyticsStream, messaging.AnalyticsGroup, "api-server-1", func(data []byte) error {
		var payload messaging.AnalyticsPayload
		if err := json.Unmarshal(data, &payload); err != nil {
			return err
		}
		analyticsService.RecordClick(context.Background(), payload.LinkID, payload.IP, payload.UserAgent, payload.Referrer)
		return nil
	})

	// Initialize controllers
	controllers := controller.NewControllers(
		authService,
		userService,
		mediaService,
		linkService,
		analyticsService,
	)

	// Create Fiber app with centralized error handler
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler(log),
	})

	// Global middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(log))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.CORSAllowOrigins},
	}))

	// Rate limiting on auth endpoints
	app.Use("/api/v1/auth", limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	}))

	// Register routes
	routes.Setup(app, controllers, jwtService, dbConn)

	// Custom Redis storage for limiter
	store := redis.New(redis.Config{
		URL: cfg.RedisURL,
	})

	// Strict rate limiting for link creation (10 per minute per IP)
	app.Post("/api/v1/links", limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		Storage:    store,
	}))

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Str("port", cfg.Port).Msg("starting server")
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
			log.Fatal().Err(err).Msg("server failed")
		}
	}()

	<-quit
	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Error().Err(err).Msg("server forced shutdown")
	}

	log.Info().Msg("server stopped")
}
