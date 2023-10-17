package handlers

import (
	"interview-follow/config"
	"interview-follow/db"
	"interview-follow/middleware"
	"interview-follow/models"
	"interview-follow/validation"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	// TODO: Remove this
	// Read all users
	user.Get("/all", middleware.DeserializeUser, GetUsers)
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func Login(c *fiber.Ctx) error {
	body := new(models.LoginRequest)
	c.BodyParser(&body)

	var user models.User
	result := db.Database.Where("email = ?", body.Email).First(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid Credentials"})
	}

	if !CheckPasswordHash(body.Password, user.HashedPassword) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Invalid Credentials"})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Signed in successfully", "data": t})
}

func SignUp(c *fiber.Ctx) error {
	body := new(models.SignupRequest)
	c.BodyParser(&body)

	// Check if the user doesn't exist
	var existingUser models.User
	result := db.Database.Where("email = ?", body.Email).First(&existingUser)

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
