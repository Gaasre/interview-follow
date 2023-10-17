package handlers

import (
	"interview-follow/db"
	"interview-follow/models"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

var pg = paginate.New()

func SetupApplicationRoutes(router fiber.Router) {
	application := router.Group("application")

	application.Get("/all", func(c *fiber.Ctx) error {
		model := db.Database.Model(&models.Application{})
		return c.JSON(pg.With(model).Request(c.Request()).Response(&[]models.Application{}))
	})
}
