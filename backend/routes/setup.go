package routes

import (
	"github.com/chirag3003/go-backend-template/controller"
	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/dto/response"
	mw "github.com/chirag3003/go-backend-template/middleware"
	"github.com/chirag3003/go-backend-template/pkg/auth"
	"github.com/gofiber/fiber/v3"
)

// Setup registers all application routes.
func Setup(app *fiber.App, controllers *controller.Controllers, jwtService *auth.JWTService, dbConn db.Connection) {
	// Health endpoints
	app.Get("/health", func(ctx fiber.Ctx) error {
		return ctx.JSON(response.OK(fiber.Map{"status": "ok"}))
	})

	app.Get("/ready", func(ctx fiber.Ctx) error {
		if err := dbConn.Ping(ctx.Context()); err != nil {
			return ctx.Status(fiber.StatusServiceUnavailable).JSON(
				response.Err("NOT_READY", "database not available"),
			)
		}
		return ctx.JSON(response.OK(fiber.Map{"status": "ready"}))
	})

	// API v1 routes
	v1 := app.Group("/api/v1")

	// Auth routes (public)
	authGroup := v1.Group("/auth")
	authGroup.Post("/login", controllers.Auth.Login)
	authGroup.Post("/register", controllers.Auth.Register)

	// User routes (protected)
	authMiddleware := mw.Auth(jwtService)
	userGroup := v1.Group("/user", authMiddleware)
	userGroup.Get("/me", controllers.User.GetMe)
	userGroup.Patch("/me", controllers.User.UpdateMe)

	// Link routes
	linksPublic := v1.Group("/links")
	linksPublic.Post("/quick", controllers.Link.Create)
	linksPublic.Get("/alias-availability", controllers.Link.AliasAvailability)

	linksProtected := v1.Group("/links", authMiddleware)
	linksProtected.Post("/", controllers.Link.Create)
	linksProtected.Get("/", controllers.Link.List)
	linksProtected.Get("/:id", controllers.Link.GetByID)
	linksProtected.Patch("/:id", controllers.Link.Update)
	linksProtected.Delete("/:id", controllers.Link.Delete)

	analyticsGroup := v1.Group("/links/:id/analytics", authMiddleware)
	analyticsGroup.Get("/summary", controllers.Analytics.Summary)
	analyticsGroup.Get("/timeseries", controllers.Analytics.TimeSeries)
	analyticsGroup.Get("/:kind", controllers.Analytics.Breakdown)

	// Media routes (protected)
	mediaGroup := v1.Group("/media", authMiddleware)
	mediaGroup.Post("/upload", controllers.Media.Upload)
}
