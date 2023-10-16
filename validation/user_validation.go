package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SignupRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string `validate:"required"`
}

func ValidateSignup(c *fiber.Ctx) error {
	var errors []*ValidationError
	body := new(SignupRequest)
	c.BodyParser(&body)

	err := Validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
