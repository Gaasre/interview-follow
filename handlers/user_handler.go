package handlers

import (
	"interview-follow/db"
	"interview-follow/models"
	"interview-follow/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	// Read all users
	user.Get("/all", func(c *fiber.Ctx) error {
		return GetUsers(c)
	})

	user.Post("/signup", validation.ValidateSignup, func(c *fiber.Ctx) error {
		return SignUp(c)
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SignUp(c *fiber.Ctx) error {
	body := new(validation.SignupRequest)
	c.BodyParser(&body)

	// Check if the user doesn't exist
	var existingUser models.User
	result := db.Database.Where("email = ?", body.Email).Find(&existingUser)

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Email already in use"})
	}

	// Create a new user
	hash, err := hashPassword(body.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	newUser := models.User{
		Id:             uuid.NewString(),
		Name:           body.Name,
		Email:          body.Email,
		HashedPassword: hash,
	}

	if err := db.Database.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User created successfully"})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.Database.Find(&users)

	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}
