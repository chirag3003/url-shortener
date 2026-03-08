package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chirag3003/go-backend-template/config"
	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/pkg/cache"
	"github.com/chirag3003/go-backend-template/pkg/logger"
	"github.com/chirag3003/go-backend-template/pkg/messaging"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
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
	redisClient := redis.NewClient(&redis.Options{Addr: cfg.RedisURL}) // Simple client for streaming
	streamManager := messaging.NewStreamManager(redisClient)

	repo := repository.NewRepository(dbConn)
	linkService := service.NewLinkService(repo.Link, cfg.BaseURL, log)

	app := fiber.New()

	app.Get("/:code", func(c fiber.Ctx) error {
		code := c.Params("code")
		ctx := c.Context()

		link, err := linkService.ResolveForRedirect(ctx, code)
		if err != nil {
			return c.Status(404).SendString("Link not found")
		}

		// Push to Analytics Stream
		go func() {
			_ = streamManager.Publish(context.Background(), messaging.AnalyticsStream, messaging.AnalyticsPayload{
				LinkID:    link.ID,
				IP:        c.IP(),
				UserAgent: string(c.Request().Header.UserAgent()),
				Referrer:  string(c.Request().Header.Referer()),
				Timestamp: time.Now(),
			})
		}()

		return c.Redirect().Status(int(link.RedirectType)).To(link.LongURL)
	})

	log.Info().Msg("Redirect Service starting on :5001")
	if err := app.Listen(":5001"); err != nil {
		log.Fatal().Err(err).Msg("server failed")
	}
}
