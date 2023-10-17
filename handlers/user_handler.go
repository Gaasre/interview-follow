package handlers

import (
	"interview-follow/config"
	"interview-follow/db"
	"interview-follow/middleware"
	"interview-follow/models"
	"interview-follow/types"
	"interview-follow/validation"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/self", middleware.DeserializeUser, GetSelf)
	user.Post("/login", validation.ValidateLogin, Login)
	user.Post("/signup", validation.ValidateSignup, SignUp)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetSelf(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	return c.Status(fiber.StatusOK).JSON(types.UserSuccess(user))
}

func Login(c *fiber.Ctx) error {
	body := new(models.LoginRequest)
	c.BodyParser(&body)

	var user models.User
	result := db.Database.Where("email = ?", body.Email).First(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(types.InvalidCredentials)
	}

	if !CheckPasswordHash(body.Password, user.HashedPassword) {
		return c.Status(fiber.StatusBadRequest).JSON(types.InvalidCredentials)
	}

	// Generate JWT Token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(types.SignInSuccess(t))
}

func SignUp(c *fiber.Ctx) error {
	body := new(models.SignupRequest)
	c.BodyParser(&body)

	// Check if the user doesn't exist
	var existingUser models.User
	result := db.Database.Where("email = ?", body.Email).First(&existingUser)

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(types.EmailAlreadyInUse)
	}

	// Create a new user
	hash, err := hashPassword(body.Password)
	if err != nil {
		return c.Status(500).JSON(types.HashError)
	}

	newUser := models.User{
		Id:             uuid.NewString(),
		Name:           body.Name,
		Email:          body.Email,
		HashedPassword: hash,
	}

	if err := db.Database.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.UserCreationFailed)
	}

	return c.Status(fiber.StatusOK).JSON(types.UserCreationSuccess)
}
