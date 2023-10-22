package handlers

import (
	"interview-follow/middleware"
	"interview-follow/openai"

	"github.com/gofiber/fiber/v2"
)

type Mail struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func SetupMailRoutes(router fiber.Router) {
	mail := router.Group("mail")

	mail.Post("/new", middleware.DeserializeUser, func(c *fiber.Ctx) error {
		body := new(Mail)
		c.BodyParser(&body)

		return c.JSON(openai.ParseEmail(body.From, body.Subject, body.Body))
	})
}
