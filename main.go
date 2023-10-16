package main

import (
	"interview-follow/db"
	"interview-follow/router"
	"interview-follow/validation"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1902"
	}

	// Connect to DB
	db.Init()

	// Init Validation
	validation.Init()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
		AllowCredentials: true,
	}))

	router.SetupRoutes(app)

	app.Listen(":" + port)
}
