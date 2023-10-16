package validation

import (
	"interview-follow/models"

	"github.com/gofiber/fiber/v2"
)

func ValidateLogin(c *fiber.Ctx) error {
	return ValidateBody(c, new(models.LoginRequest))
}

func ValidateSignup(c *fiber.Ctx) error {
	return ValidateBody(c, new(models.SignupRequest))
}
