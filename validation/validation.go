package validation

import (
	"interview-follow/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func ErrorMap(err error, errors []*ValidationError) []*ValidationError {
	for _, err := range err.(validator.ValidationErrors) {
		var el ValidationError
		el.Field = err.Field()
		el.Tag = err.Tag()
		el.Value = err.Param()
		errors = append(errors, &el)
	}
	return errors
}

func ValidateBody[T any](c *fiber.Ctx, payload T) error {
	var errors []*ValidationError
	c.BodyParser(payload)

	err := validate.Struct(payload)
	if err != nil {
		errors = ErrorMap(err, errors)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func ValidateLogin(c *fiber.Ctx) error {
	return ValidateBody(c, new(models.LoginRequest))
}

func ValidateSignup(c *fiber.Ctx) error {
	return ValidateBody(c, new(models.SignupRequest))
}

func ValidateNewApplication(c *fiber.Ctx) error {
	return ValidateBody(c, new(models.ApplicationRequest))
}

func ValidateEditApplication(c *fiber.Ctx) error {
	return ValidateBody(c, new(models.ApplicationRequest))
}
