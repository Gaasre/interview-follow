package middleware

import (
	"fmt"
	"interview-follow/config"
	"interview-follow/db"
	"interview-follow/models"
	"interview-follow/types"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {
	var token string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		token = strings.TrimPrefix(authorization, "Bearer ")
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(types.Unauthorized)
	}

	tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte(config.Config("SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(types.InvalidateToken)
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(types.InvalidTokenClaim)
	}

	var user models.User
	db.Database.First(&user, "id = ?", fmt.Sprint(claims["user_id"]))

	if user.Id != claims["user_id"] {
		return c.Status(fiber.StatusForbidden).JSON(types.InvalidUser)
	}

	c.Locals("user", models.FilterPassword(user))
	return c.Next()
}
