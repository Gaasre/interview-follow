package router

import (
	route "interview-follow/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	route.SetupMailRoutes(api)
	route.SetupUserRoutes(api)
	route.SetupApplicationRoutes(api)
	route.SetupInterviewRoutes(api)
}
