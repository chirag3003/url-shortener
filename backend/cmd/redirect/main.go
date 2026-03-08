package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chirag3003/go-backend-template/config"
	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/middleware"
	"github.com/chirag3003/go-backend-template/pkg/cache"
	"github.com/chirag3003/go-backend-template/pkg/logger"
	"github.com/chirag3003/go-backend-template/pkg/messaging"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	log := logger.New(cfg.LogLevel)
	dbConn, err := db.ConnectPostgres(cfg, log)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to PostgreSQL")
	}
	defer dbConn.Close()

	if err := cache.Init(cfg.RedisURL); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize Redis")
	}
	opts, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse redis URL")
	}
	redisClient := redis.NewClient(opts) // Simple client for streaming
	streamManager := messaging.NewStreamManager(redisClient)

	repo := repository.NewRepository(dbConn)
	linkService := service.NewLinkService(repo.Link, cfg.BaseURL, log)

	app := fiber.New(fiber.Config{
		ErrorHandler:   middleware.ErrorHandler(log),
		ReadBufferSize: 16384, // Increase to 16KB to handle large headers/cookies
	})

	// Global middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(log))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.CORSAllowOrigins},
	}))

	app.Get("/:code", func(c fiber.Ctx) error {
		code := c.Params("code")
		ctx := c.Context()
		log.Debug().Str("code", code).Msg("resolving link")

		link, err := linkService.ResolveForRedirect(ctx, code)
		if err != nil {
			log.Error().Err(err).Str("code", code).Msg("link not found")
			return c.Status(404).SendString("Link not found")
		}

		// Capture request data BEFORE starting the goroutine
		// Fiber's Ctx is pooled and reused; accessing it in a goroutine after the handler returns is unsafe.
		ip := c.IP()
		userAgent := string(c.Request().Header.UserAgent())
		referer := string(c.Request().Header.Referer())
		linkID := link.ID
		redirectType := int(link.RedirectType)
		longURL := link.LongURL

		// Push to Analytics Stream asynchronously
		go func(lID int64, remoteIP, ua, ref string) {
			err := streamManager.Publish(context.Background(), messaging.AnalyticsStream, messaging.AnalyticsPayload{
				LinkID:    lID,
				IP:        remoteIP,
				UserAgent: ua,
				Referrer:  ref,
				Timestamp: time.Now(),
			})
			if err != nil {
				log.Error().Err(err).Int64("linkID", lID).Msg("failed to publish analytics")
			}
		}(linkID, ip, userAgent, referer)

		return c.Redirect().Status(redirectType).To(longURL)
	})

	log.Info().Msg("Redirect Service starting on :5001")
	if err := app.Listen(":5001"); err != nil {
		log.Fatal().Err(err).Msg("server failed")
	}
}
